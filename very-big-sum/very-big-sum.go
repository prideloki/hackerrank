package main

import (
	"fmt"
	"math/big"
)

func main() {
	var length int
	sum := new(big.Int)

	fmt.Scanf("%v", &length)
	for i := 0; i < length; i++ {
		var num big.Int
		fmt.Scanf("%v", &num)
		sum.Add(sum, &num)
	}
	fmt.Println(sum.String())
}
