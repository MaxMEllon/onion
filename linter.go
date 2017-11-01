package onion

import (
	"github.com/maxmellon/onion/rules"
	"github.com/maxmellon/onion/types"
)

func Linting(tml []types.Tml) []types.Status {
	var result []types.Status
	for _, v := range tml {
		r := rules.PunctuationMark(v)
		if r.Code != types.S {
			result = append(result, r)
		}

		r = rules.LessThan80Characters(v)
		if r.Code != types.S {
			result = append(result, r)
		}
	}
	return result
}
