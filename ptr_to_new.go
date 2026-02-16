package ptrtonew

import (
	"fmt"
	"go/ast"
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
		ptrName, imported := ptrImportExists(file)

		if !imported {
			continue
		}

		ast.Inspect(file, func(n ast.Node) bool {
			value, casting := getPtrToSel(n, ptrName)
			if value == nil {
				return true
			}

			value = astcopy.Expr(value)
			if casting != nil {
				casting = astcopy.Expr(casting)

				value = &ast.CallExpr{Fun: casting, Args: []ast.Expr{value}}
			}

			suggestion := &ast.CallExpr{Fun: ast.NewIdent("new"), Args: []ast.Expr{value}}
			suggestionText := format.Format(suggestion)

			pass.Report(analysis.Diagnostic{
				Pos:      n.Pos(),
				End:      n.End(),
				Category: "modernize",
				Message:  fmt.Sprintf(`use the "new" operator instead of %s.To; %s`, ptrName, suggestionText),
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

func getPtrToSel(n ast.Node, ptrName string) (value ast.Expr, casting ast.Expr) {
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

	case *ast.IndexExpr:
		if sel, ok := fun.X.(*ast.SelectorExpr); ok && isPtrTo(sel, ptrName) {
			return value, fun.Index
		}

		return nil, nil
	}

	return nil, nil
}

func isPtrTo(fun *ast.SelectorExpr, ptrName string) bool {
	if fun.Sel.Name != "To" {
		return false
	}

	id, ok := fun.X.(*ast.Ident)
	return ok && id.Name == ptrName
}

func ptrImportExists(file *ast.File) (string, bool) {
	idx := slices.IndexFunc(file.Imports, func(imp *ast.ImportSpec) bool {
		return imp.Path.Value == `"k8s.io/utils/ptr"`
	})

	if idx == -1 {
		return "", false
	}

	imp := file.Imports[idx]
	if imp.Name == nil {
		return "ptr", true
	}

	return imp.Name.Name, true
}
