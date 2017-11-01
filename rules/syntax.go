package rules

import (
	"github.com/maxmellon/onion/types"
)

func PunctuationMark(m types.Tml) types.Status {
	if m.Category == "PLANE_TEXT" {
		for i, v := range m.Statement {
			if v == '．' || v == '，' || v == '.' || v == ',' {
				return types.Status{
					Line:     m.Line,
					Column:   i + 1,
					Message:  "`．``，` を使わず `、``。` を使うべきである",
					RuleName: "PunctuationMark",
					Code:     types.C,
				}
			}
		}
	}
	return types.Status{Code: types.S}
}

func LessThan80Characters(m types.Tml) types.Status {
	for i, _ := range m.Statement {
		if i >= 80 {
			return types.Status{
				Line:     m.Line,
				Column:   i + 1,
				Message:  "1行は80文字以内にするべきである",
				RuleName: "LessThan80Characters",
				Code:     types.W,
			}
		}
	}
	return types.Status{Code: types.S}
}
