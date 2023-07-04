package bigint

import (
	"errors"
)

var ErrorBadInput = errors.New("bad input")

type Bigint struct {
	value string
}

func NewInt(num string) (Bigint, error) {
	err := validate(num)
	if err != nil {
		return Bigint{}, err
	}
	num = simplify(num)
	return Bigint{value: num}, nil
}

// func (z *Bigint) Result() string {
// 	return z.value
// }
