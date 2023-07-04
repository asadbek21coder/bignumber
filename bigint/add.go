package bigint

import (
	"fmt"
	"strconv"
)

func Add(a, b Bigint) Bigint {
	// a va b ni bir biriga qo'shib qaytarishi kk
	// fmt.Println(a.value, b.value)
	var big, small string

	if len(a.value) >= len(b.value) {
		big, small = a.value, b.value
	} else {
		big, small = b.value, a.value
	}
	fmt.Println(add(big, small))
	fmt.Println(subtract(big, small))
	return Bigint{}
}

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
	return ""
}

// 1
// 463474583472
// 800000236462

// 463474583472
//       236462
