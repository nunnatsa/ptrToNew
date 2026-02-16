package formatter

import (
	"bytes"
	"go/ast"
	"go/printer"
	"go/token"
)

type GoFmtFormatter struct {
	fset *token.FileSet
}

func NewGoFmtFormatter(fset *token.FileSet) *GoFmtFormatter {
	return &GoFmtFormatter{fset: fset}
}

func (f GoFmtFormatter) Format(exp ast.Expr) (string, error) {
	var buf bytes.Buffer
	err := printer.Fprint(&buf, f.fset, exp)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}
