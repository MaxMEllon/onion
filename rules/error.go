package rules

import (
	"github.com/TomiLabo/onion/types"
	"github.com/TomiLabo/tmngparser/ast"
	"github.com/TomiLabo/tmngparser/token"
)

func ParseError(m ast.Tml) types.Status {
	if m.Token == token.Error {
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
