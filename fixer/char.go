package fixer

import (
	"regexp"

	"github.com/TomiLabo/tmngparser/ast"
)

func ReplaceDotAndConma(node ast.Tml) ast.Tml {
	exp1 := regexp.MustCompile("．")
	exp2 := regexp.MustCompile("，")
	exp3 := regexp.MustCompile("　")
	exp4 := regexp.MustCompile("◯")
	statement := exp1.ReplaceAllString(node.Statement, "。")
	statement = exp2.ReplaceAllString(statement, "、")
	statement = exp3.ReplaceAllString(statement, " ")
	statement = exp4.ReplaceAllString(statement, "○")
	return ast.New(
		node.Line,
		node.Token,
		statement,
		node.Start,
		node.End,
	)
}
