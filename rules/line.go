package rules

import (
	"github.com/maxmellon/onion/types"
)

const MAX_LENGTH_OF_LINE = 75

func LineLength(m types.Tml) types.Status {
	for i, _ := range m.Statement {
		if i >= MAX_LENGTH_OF_LINE {
			return types.Status{
				Line:     m.Line,
				Column:   i + 1,
				Message:  "1行は75文字以内にするべき",
				RuleName: "Line/Length",
				Code:     types.W,
			}
		}
	}
	return types.Status{Code: types.S}
}
