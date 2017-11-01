package onion

import (
	"regexp"
)

type TmngMail struct {
	category  string
	line      int
	statement string
}

func check_regexp(reg, str string) bool {
	return regexp.MustCompile(reg).Match([]byte(str))
}

func generate(i int, c string, s string) TmngMail {
	return TmngMail{
		line:      i,
		category:  c,
		statement: s,
	}
}

func Parse(data []string) []TmngMail {
	var result []TmngMail
	for i, l := range data {
		if check_regexp(`^■`, l) {
			result = append(result, generate(i, "HEADER", l))
		} else if check_regexp(`^●`, l) {
			result = append(result, generate(i, "TITLE", l))
		} else if check_regexp(`^◎`, l) {
			result = append(result, generate(i, "SUB_TITLE", l))
		} else if check_regexp(`^○`, l) {
			result = append(result, generate(i, "LIST_ITEM", l))
		} else if check_regexp(`^\n|^\r`, l) {
			result = append(result, generate(i, "EMPTY", l))
		} else {
			result = append(result, generate(i, "PLANE_TEXT", l))
		}
	}
	return result
}
