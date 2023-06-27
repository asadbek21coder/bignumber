package bigint

import (
	"errors"
)

type Bigint struct {
	value string
}

func (z *Bigint) Result() string {
	return z.value
}

var ErrorBadInput = errors.New("bad input")

func Validate(num string) error {

	// faqat raqamdan iborat bo'lishi kk
	// noto'g'ri bo'lsa ErrorBadInput qaytarishi kk

	return nil
}

func Simplify(num string) string {
	// oldida + bo'sa kesish kk
	// Ko'rinishini soddalashtirishi kk
	return num
}

func NewInt(num string) (Bigint, error) {
	err := Validate(num)
	if err != nil {
		return Bigint{}, err
	}
	num = Simplify(num)
	return Bigint{value: num}, nil
}

func Add(a, b Bigint) Bigint {
	// a va b ni bir biriga qo'shib qaytarishi kk
	return Bigint{value: a.value + b.value}
}

func Subtract(a, b Bigint) Bigint {

	// a va b ni ayrimasini hisoblashi kk

	return Bigint{}
}

func Multiply(a, b Bigint) Bigint {
	// a va b ni ko'paytirib berishi kk
	return Bigint{}
}

// ***
func Divide(a, b Bigint) Bigint {
	// a ni b ga bo'lib qaytaradi

	return Bigint{}
}

func (b *Bigint) Abs() Bigint {
	// Sonning modulini qaytaradi
	return *b
}
