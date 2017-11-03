package onion

import (
	"regexp"

	"github.com/maxmellon/onion/types"
)

func match(reg, str string) bool {
	return regexp.MustCompile(reg).Match([]byte(str))
}

func generate(i int, c string, s string) types.Tml {
	return types.Tml{
		Line:      i + 1,
		Category:  c,
		Statement: s,
	}
}

func Parse(data []string) []types.Tml {
	var result []types.Tml
	for i, l := range data {
		if match(`^■`, l) {
			result = append(result, generate(i, "HEADER", l))
		} else if match(`^●`, l) {
			result = append(result, generate(i, "TITLE", l))
		} else if match(`^◎`, l) {
			result = append(result, generate(i, "SUB_TITLE", l))
		} else if match(`^○`, l) {
			result = append(result, generate(i, "LIST_ITEM", l))
		} else if match(`^\n|^\r`, l) {
			result = append(result, generate(i, "EMPTY", l))
		} else {
			result = append(result, generate(i, "PLANE_TEXT", l))
		}
	}
	return result
}
