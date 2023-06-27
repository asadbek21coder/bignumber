package main

import (
	"fmt"

	"github.com/asadbek21coder/bignumber/bigint"
)

func main() {
	// fmt.Println("Hello world")
	a, err := bigint.NewInt("124925489655546454")
	if err != nil {
		fmt.Println(err)
	}

	b, err := bigint.NewInt("468216132684621684621684")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(bigint.Add(a, b))

}
