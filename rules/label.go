package rules

import (
	"fmt"
	"github.com/maxmellon/onion/types"
	"regexp"
)

func contains(s []string, t string) bool {
	for _, v := range s {
		regExp := fmt.Sprintf("%s", v)
		reg := regexp.MustCompile(regExp)
		if reg.MatchString(t) {
			return true
		}
	}
	return false
}

var labels = []string{"新規", "継続", "完了"}

func LabelStyle(m types.Tml) types.Status {
	if m.Category != "LIST_ITEM" {
		return types.Status{Code: types.S}
	}

	rep := regexp.MustCompile(`○ (\S+) \[(\d+)\]`)
	strList := rep.FindAllString(m.Statement, 1)
	for i, v := range strList {
		if !contains(labels, v) {
			return types.Status{
				Line:     m.Line,
				Column:   i + 1,
				Message:  "[新規], [継続], [完了] のいずれかであるべき",
				RuleName: "LabelStyle",
				Code:     types.C,
			}
		}
	}
	return types.Status{Code: types.S}
}
