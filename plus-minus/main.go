package main

import (
	"fmt"
)

func main() {
	var length, input, pos, neg, zero int
	var fracPos, fracNeg, fracZero float64
	fmt.Scanf("%v", &length)

	for i := 0; i < length; i++ {
		fmt.Scanf("%v", &input)
		if input == 0 {
			zero++
		} else if input < 0 {
			neg++
		} else {
			pos++
		}
	}
	fracPos = float64(pos) / float64(length)
	fracNeg = float64(neg) / float64(length)
	fracZero = float64(zero) / float64(length)
	fmt.Printf("%f\n", fracPos)
	fmt.Printf("%f\n", fracNeg)
	fmt.Printf("%f\n", fracZero)
}
