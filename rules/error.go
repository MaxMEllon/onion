package rules

import (
	"github.com/maxmellon/onion/types"
)

func ParseError(m types.Tml) types.Status {
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
