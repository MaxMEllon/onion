package rules

import (
	"regexp"

	"github.com/TomiLabo/onion/types"
	"github.com/TomiLabo/tmngparser/ast"
	"github.com/TomiLabo/tmngparser/token"
)

func PunctuationMark(m ast.Tml) types.Status {
	if m.Token != token.PlaneText {
		return types.Status{Code: types.S}
	}
	for i, v := range m.Statement {
		if v == '．' || v == '，' {
			return types.Status{
				Line:     m.Line,
				Column:   i + 1,
				Message:  "`．``，` を使わず `、``。` を使うべき",
				RuleName: "Char/PunctuationMark",
				Code:     types.C,
			}
		}
	}
	return types.Status{Code: types.S}
}

func VariantCircleCharacter(m ast.Tml) types.Status {
	reg := regexp.MustCompile("◯")
	loc := reg.FindStringIndex(m.Statement)
	if loc == nil {
		return types.Status{Code: types.S}
	}
	return types.Status{
		Line:     m.Line,
		Column:   loc[1],
		Message:  "◯ ではなく，○ を使うべき",
		RuleName: "Char/VariantCircleCharacter",
		Code:     types.C,
	}
}

func AmbiWidthSpace(m ast.Tml) types.Status {
	reg := regexp.MustCompile("　")
	loc := reg.FindStringIndex(m.Statement)
	if loc == nil {
		return types.Status{Code: types.S}
	}
	return types.Status{
		Line:     m.Line,
		Column:   loc[1],
		Message:  "全角スペースを使うべきではない",
		RuleName: "Char/AmbiWidthSpace",
		Code:     types.C,
	}
}
