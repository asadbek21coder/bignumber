package bigint

func Add(a, b Bigint) Bigint {
	var big, small string
	var sign bool

	if string(a.value[0]) == "-" && string(b.value[0]) == "-" {
		big, small, _ = bigSmall(a.value, b.value)
		return Bigint{value: "-" + add(big, small)}
	} else if string(a.value[0]) == "-" || string(b.value[0]) == "-" {
		big, small, sign = bigSmall(a.value, b.value)
		if !sign {
			return Bigint{value: "-" + subtract(big, small)}
		}
		return Bigint{value: subtract(big, small)}
	}

	big, small, _ = bigSmall(a.value, b.value)
	return Bigint{value: add(big, small)}

}
