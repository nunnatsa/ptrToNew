package ptrtonew

import (
	"fmt"
	"go/ast"
	"go/types"
	"os"
	"ptrtonew/formatter"
	"slices"

	"github.com/go-toolsmith/astcopy"
	"golang.org/x/tools/go/analysis"
)

func NewAnalyzer() *analysis.Analyzer {
	return &analysis.Analyzer{
		Name: "ptrToNew",
		Doc: `In Go 1.26, the built-in new function now allows its operand to be an expression,
specifying the initial value of the variable. This makes k8s.io/utils/ptr.To unnecessary.
This tool finds all usages of ptr.To and suggests replacing them with the new function.`,
		Run: run,
	}
}

func run(pass *analysis.Pass) (any, error) {
	format := formatter.NewGoFmtFormatter(pass.Fset)
	for _, file := range pass.Files {
		ptrName, ptrImported := importExists(file, `"k8s.io/utils/ptr"`, "ptr")
		pointerName, pointerImported := importExists(file, `"k8s.io/utils/pointer"`, "pointer")

		if !ptrImported && !pointerImported {
			continue
		}

		ast.Inspect(file, func(n ast.Node) bool {
			value, casting := getPtrToSel(pass, n, ptrName, pointerName)
			if value == nil {
				return true
			}

			value = astcopy.Expr(value)
			if casting != nil {
				casting = astcopy.Expr(casting)

				value = &ast.CallExpr{Fun: casting, Args: []ast.Expr{value}}
			}

			pos := pass.Fset.Position(n.Pos()).String()

			origText, err := format.Format(n.(ast.Expr))
			if err != nil {
				fmt.Fprintf(os.Stderr, "cannot generate suggestion text in %s; %v\n", pos, err)
			}

			suggestion := &ast.CallExpr{Fun: ast.NewIdent("new"), Args: []ast.Expr{value}}
			suggestionText, err := format.Format(suggestion)
			if err != nil {
				fmt.Fprintf(os.Stderr, "cannot generate suggestion text in %s; %v\n", pos, err)
				pass.Reportf(n.Pos(), `replace %s with the "new()" built-in function`, origText)
				return false
			}

			pass.Report(analysis.Diagnostic{
				Pos:      n.Pos(),
				End:      n.End(),
				Category: "modernize",
				Message:  fmt.Sprintf(`replace %s with the "new()" built-in function: %s`, origText, suggestionText),
				SuggestedFixes: []analysis.SuggestedFix{
					{
						Message: "Replace with new()",
						TextEdits: []analysis.TextEdit{
							{
								Pos:     n.Pos(),
								End:     n.End(),
								NewText: []byte(suggestionText),
							},
						},
					},
				},
			})

			return false
		})
	}

	return nil, nil
}

func getPtrToSel(pass *analysis.Pass, n ast.Node, ptrName, pointerName string) (value ast.Expr, casting ast.Expr) {
	call, ok := n.(*ast.CallExpr)
	if !ok {
		return nil, nil
	}

	if len(call.Args) != 1 {
		return nil, nil
	}

	value = call.Args[0]

	switch fun := call.Fun.(type) {
	case *ast.SelectorExpr:
		if isPtrTo(fun, ptrName) {
			return value, nil
		}

		return handlePointer(pass, fun, pointerName, value)

	case *ast.IndexExpr:
		if sel, ok := fun.X.(*ast.SelectorExpr); ok && isPtrTo(sel, ptrName) {
			return value, fun.Index
		}

		return nil, nil
	}

	return nil, nil
}

func importExists(file *ast.File, importPath, defaultName string) (string, bool) {
	idx := slices.IndexFunc(file.Imports, func(imp *ast.ImportSpec) bool {
		return imp.Path.Value == importPath
	})

	if idx == -1 {
		return "", false
	}

	imp := file.Imports[idx]
	if imp.Name == nil {
		return defaultName, true
	}

	return imp.Name.Name, true
}

func isPtrTo(fun *ast.SelectorExpr, ptrName string) bool {
	if fun.Sel.Name != "To" {
		return false
	}

	id, ok := fun.X.(*ast.Ident)
	return ok && id.Name == ptrName
}

func isPointer(fun *ast.SelectorExpr, pointerName string) (itIsAPointer bool, casting string) {
	id, ok := fun.X.(*ast.Ident)
	if !ok || id.Name != pointerName {
		return false, ""
	}

	return isPointerName(fun.Sel.Name)
}

func isPointerName(name string) (IsAPointerName bool, casting string) {
	switch name {
	case "String", "Bool", "Duration":
		casting = ""

	case "Int":
		casting = "int"
	case "Int32":
		casting = "int32"
	case "Uint":
		casting = "uint"
	case "Uint32":
		casting = "uint32"
	case "Int64":
		casting = "int64"
	case "Uint64":
		casting = "uint64"
	case "Float64":
		casting = "float64"
	case "Float32":
		casting = "float32"

	default:
		return false, ""
	}

	return true, casting
}

func getCastingForUntypedConst(pass *analysis.Pass, id *ast.Ident, castingName string) *ast.Ident {
	objType := pass.TypesInfo.ObjectOf(id)
	c, ok := objType.(*types.Const)
	if !ok {
		return nil
	}

	basicType, ok := c.Type().(*types.Basic)
	if !ok {
		return nil
	}

	kind := basicType.Kind()

	if kind != types.UntypedInt && kind != types.UntypedFloat {
		return nil
	}

	if (castingName == "int" && kind == types.UntypedInt) ||
		(castingName == "float64" && kind == types.UntypedFloat) {
		return nil
	}

	return ast.NewIdent(castingName)
}

func handlePointer(pass *analysis.Pass, fun *ast.SelectorExpr, pointerName string, value ast.Expr) (ast.Expr, ast.Expr) {
	isAPointer, castingName := isPointer(fun, pointerName)
	if isAPointer {
		if castingName == "" {
			return value, nil
		}

		switch t := value.(type) {
		case *ast.BasicLit:
			if t.Kind.IsLiteral() {
				if (castingName == "int" && t.Kind.String() != "FLOAT") ||
					(castingName == "float64" && t.Kind.String() != "INT") {
					return value, nil
				}

				return value, ast.NewIdent(castingName)
			}

		case *ast.Ident:
			if castingIdent := getCastingForUntypedConst(pass, t, castingName); castingIdent != nil {
				return value, castingIdent
			}

		case *ast.SelectorExpr:
			if castingIdent := getCastingForUntypedConst(pass, t.Sel, castingName); castingIdent != nil {
				return value, castingIdent
			}
		}

		return value, nil
	}

	return nil, nil
}
