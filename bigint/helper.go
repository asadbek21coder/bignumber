package bigint

import (
	"fmt"
	"strconv"
	"strings"
)

func add(big, small string) string {
	var res string
	for len(big)-len(small) > 0 {
		small = "0" + small
	}
	var flag int
	for i := len(big) - 1; i >= 0; i-- {
		a1, _ := strconv.Atoi(string(big[i]))
		a2, _ := strconv.Atoi(string(small[i]))
		res = fmt.Sprint((a1+a2+flag)%10) + res
		flag = (a1 + a2 + flag) / 10

	}
	if flag > 0 {
		res = fmt.Sprint(flag) + res
	}
	return res
}

func subtract(big, small string) string {
	var res string
	for len(big)-len(small) > 0 {
		small = "0" + small
	}
	var flag int
	for i := len(big) - 1; i >= 0; i-- {
		a1, _ := strconv.Atoi(string(big[i]))
		a2, _ := strconv.Atoi(string(small[i]))
		if a1-flag >= a2 {
			res = fmt.Sprint(a1-flag-a2) + res
			flag = 0
		} else {
			res = fmt.Sprint(10+a1-flag-a2) + res
			flag = 1
		}
	}
	return res
}

func validate(num string) error {
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

func simplify(num string) string {
	for (string(num[0]) == "0" && len(num) != 1) || (string(num[0]) == "+") {
		num = num[1:]
	}
	for string(num[0]) == "-" && len(num) != 1 && string(num[1]) == "0" {
		num = num[:1] + num[2:]
	}
	return num
}

func bigSmall(a, b string) (string, string, bool) {
	var signA, signB bool = true, true
	if string(a[0]) == "-" {
		signA = false
		a = a[1:]
	}
	if string(b[0]) == "-" {
		b = b[1:]
		signB = false
	}

	if len(a) > len(b) {
		return a, b, signA
	} else if len(a) < len(b) {
		return b, a, signB
	}
	for i := 0; i < len(a); i++ {
		if a[i] == b[i] {
			continue
		}
		if a[i] > b[i] {
			return a, b, signA
		}
		if a[i] < b[i] {
			return b, a, signB
		}
	}

	return a, b, signA
}
