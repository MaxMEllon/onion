package rules

import (
	"regexp"

	"github.com/maxmellon/onion/types"
)

func PunctuationMark(m types.Tml) types.Status {
	if m.Category != "PLANE_TEXT" {
		return types.Status{Code: types.S}
	}
	for i, v := range m.Statement {
		if v == '．' || v == '，' {
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

func AmbiWidthSpace(m types.Tml) types.Status {
	reg := regexp.MustCompile("　")
	loc := reg.FindStringIndex(m.Statement)
	if loc == nil {
		return types.Status{Code: types.S}
	}
	return types.Status{
		Line:     m.Line,
		Column:   loc[1],
		Message:  "全角スペースを使うべきではない",
		RuleName: "AmbiWidthSpace",
		Code:     types.C,
	}
}

const MAX_LENGTH_OF_LINE = 75

func LessThan75Characters(m types.Tml) types.Status {
	for i, _ := range m.Statement {
		if i >= MAX_LENGTH_OF_LINE {
			return types.Status{
				Line:     m.Line,
				Column:   i + 1,
				Message:  "1行は75文字以内にするべき",
				RuleName: "LessThan75Characters",
				Code:     types.W,
			}
		}
	}
	return types.Status{Code: types.S}
}
