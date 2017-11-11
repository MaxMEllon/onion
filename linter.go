package onion

import (
	"github.com/TomiLabo/onion/rules"
	"github.com/TomiLabo/onion/types"
	"github.com/TomiLabo/tmngparser/ast"
)

func Linting(tml []ast.Tml) []types.Status {
	var result []types.Status
	for _, v := range tml {
		r := rules.ParseError(v)
		if r.Code != types.S {
			result = append(result, r)
		}

		r = rules.PunctuationMark(v)
		if r.Code != types.S {
			result = append(result, r)
		}

		r = rules.VariantCircleCharacter(v)
		if r.Code != types.S {
			result = append(result, r)
		}

		r = rules.AmbiWidthSpace(v)
		if r.Code != types.S {
			result = append(result, r)
		}

		r = rules.LineLength(v)
		if r.Code != types.S {
			result = append(result, r)
		}

		r = rules.LabelType(v)
		if r.Code != types.S {
			result = append(result, r)
		}
	}
	return result
}
