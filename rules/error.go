package rules

import (
	"github.com/TomiLabo/onion/types"
	"github.com/TomiLabo/tmngparser/ast"
)

func ParseError(m ast.Tml) types.Status {
	if m.Category == "ERROR" {
		return types.Status{
			Line:     m.Line,
			Column:   1,
			Message:  "行頭にスペースが入っています",
			RuleName: "ParseError",
			Code:     types.W,
		}
	}
	return types.Status{Code: types.S}
}
