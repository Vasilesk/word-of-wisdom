package typeutils

import "fmt"

func NewStringer(str string) fmt.Stringer {
	return stringHolder(str)
}

type stringHolder string

func (s stringHolder) String() string {
	return string(s)
}
