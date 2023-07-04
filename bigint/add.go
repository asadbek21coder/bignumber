package bigint

import (
	"fmt"
)

func Add(a, b Bigint) Bigint {
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

