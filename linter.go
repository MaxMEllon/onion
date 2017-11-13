package onion

import (
	"github.com/TomiLabo/onion/rules"
	"github.com/TomiLabo/onion/types"
	"github.com/TomiLabo/tmngparser/ast"
)

func appendResult(results []types.Status, r types.Status) []types.Status {
	if r.Code != types.S {
		results = append(results, r)
	}
	return results
}

func Linting(tml []ast.Tml) []types.Status {
	var result []types.Status
	for _, v := range tml {
		r := rules.ParseError(v)
		result = appendResult(result, r)

		r = rules.PunctuationMark(v)
		result = appendResult(result, r)

		r = rules.VariantCircleCharacter(v)
		result = appendResult(result, r)

		r = rules.AmbiWidthSpace(v)
		result = appendResult(result, r)

		r = rules.LineLength(v)
		result = appendResult(result, r)

		r = rules.LabelType(v)
		result = appendResult(result, r)
	}
	return result
}
