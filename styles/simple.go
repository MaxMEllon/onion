package styles

import (
	"fmt"
	"github.com/maxmellon/onion/types"
)

func SimplePrintResult(r []types.Status) {
	for _, v := range r {
		fmt.Printf("%s: %3d:%3d: [%s]: %s\n", v.Code, v.Line, v.Column, v.RuleName, v.Message)
	}
}
