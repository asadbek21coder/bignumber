package main

import (
	"fmt"
	"log"
	"time"

	"github.com/asadbek21coder/bignumber/bigint"
)

func main() {

	var one, two string
	fmt.Scan(&one)
	fmt.Scan(&two)

	start := time.Now()

	a, err := bigint.NewInt(one)
	if err != nil {
		fmt.Println(err)
	}

	b, err := bigint.NewInt(two)
	if err != nil {
		fmt.Println(err)
	}

	// fmt.Println(a, b)
	fmt.Println(bigint.Add(a, b))

	elapsed := time.Since(start)
	log.Printf("it took %s", elapsed)

}
