package onion

import (
	"github.com/TomiLabo/onion/fixer"
	"github.com/TomiLabo/tmngparser/ast"
)

func Fixing(tree []ast.Tml) []ast.Tml {
	var newTree []ast.Tml
	newTree = make([]ast.Tml, len(tree))
	copy(newTree, tree)
	for i, t := range tree {
		r := fixer.ReplaceDotAndConma(t)
		r = fixer.TrailingSpace(r)
		newTree[i] = r
	}
	return newTree
}
