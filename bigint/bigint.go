package bigint

import (
	"errors"
	"strings"
)

var ErrorBadInput = errors.New("bad input")

type Bigint struct {
	value string
}

// func (z *Bigint) Result() string {
// 	return z.value
// }

func Validate(num string) error {
	// faqat raqamdan iborat bo'lishi kk
	// noto'g'ri bo'lsa ErrorBadInput qaytarishi kk
	allowed := "1234567890"

	if !strings.Contains("+-"+allowed, string(num[0])) {
		return ErrorBadInput
	}

	for i := 1; i < len(num); i++ {
		if !strings.Contains(allowed, string(num[i])) {
			return ErrorBadInput
		}
	}
	return nil
}

func Simplify(num string) string {
	for (string(num[0]) == "0" && len(num) != 1) || (string(num[0]) == "+") {
		num = num[1:]
	}
	for string(num[0]) == "-" && len(num) != 1 && string(num[1]) == "0" {
		num = num[:1] + num[2:]
	}
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
