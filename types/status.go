package types

type StatusCode int

func (i StatusCode) String() string {
	switch i {
	case C:
		return "C"
	case W:
		return "W"
	case S:
		return "S"
	default:
		return ""
	}
}

const (
	C StatusCode = iota
	W
	S
)

type Status struct {
	Line     int
	Column   int
	Message  string
	RuleName string
	Code     StatusCode
}
