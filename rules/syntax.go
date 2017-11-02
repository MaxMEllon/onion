package rules

import (
	"github.com/maxmellon/onion/types"
)

func PunctuationMark(m types.Tml) types.Status {
	if m.Category != "PLANE_TEXT" {
		return types.Status{Code: types.S}
	}
	for i, v := range m.Statement {
		if v == '．' || v == '，' || v == '.' || v == ',' {
			return types.Status{
				Line:     m.Line,
				Column:   i + 1,
				Message:  "`．``，` を使わず `、``。` を使うべき",
				RuleName: "PunctuationMark",
				Code:     types.C,
			}
		}
	}
	return types.Status{Code: types.S}
}

const MAX_LENGTH_OF_LINE = 72

func LessThan72Characters(m types.Tml) types.Status {
	for i, _ := range m.Statement {
		if i >= MAX_LENGTH_OF_LINE {
			return types.Status{
				Line:     m.Line,
				Column:   i + 1,
				Message:  "1行は72文字以内にするべき",
				RuleName: "LessThan72Characters",
				Code:     types.W,
			}
		}
	}
	return types.Status{Code: types.S}
}
