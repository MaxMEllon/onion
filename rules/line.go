package rules

import (
	"github.com/TomiLabo/onion/types"
	"github.com/TomiLabo/tmngparser/ast"
)

const MAX_LENGTH_OF_LINE = 75

func LineLength(m ast.Tml) types.Status {
	if m.End >= MAX_LENGTH_OF_LINE {
		return types.Status{
			Line:     m.Line,
			Column:   m.End,
			Message:  "1行は75文字以内にするべき",
			RuleName: "Line/Length",
			Code:     types.W,
		}
	}
	return types.Status{Code: types.S}
}
