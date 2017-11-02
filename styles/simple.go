package styles

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/maxmellon/onion/types"
)

func SimplePrintResult(filename string, r []types.Status) {
	if len(r) == 0 {
		return
	}
	color.Yellow("== %s ==\n", filename)
	for _, v := range r {
		if v.Code == types.C {
			y := color.New(color.FgYellow, color.Bold)
			y.Printf("%s", v.Code)
		}
		if v.Code == types.W {
			y := color.New(color.FgHiBlue, color.Bold)
			y.Printf("%s", v.Code)
		}
		fmt.Printf(": %3d:%3d: [%s]: %s\n", v.Line, v.Column, v.RuleName, v.Message)
	}
}
