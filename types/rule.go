package types

type Rule interface {
	Check(m Tml) Status
}
