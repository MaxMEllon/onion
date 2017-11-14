package fixer

import (
	"regexp"

	"github.com/TomiLabo/tmngparser/ast"
	"github.com/TomiLabo/tmngparser/token"
)

func TrailingSpace(node ast.Tml) ast.Tml {
	exp1 := regexp.MustCompile(`[\t\ ]+$`)
	statement := exp1.ReplaceAllLiteralString(node.Statement, "")
	if node.Token != token.PlaneText {
		exp2 := regexp.MustCompile(`^[\t\ ]+`)
		statement = exp2.ReplaceAllLiteralString(statement, "")
	}
	return ast.New(
		node.Line,
		node.Token,
		statement,
		node.Start,
		node.End,
	)
}
