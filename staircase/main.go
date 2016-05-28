package main

import (
	"fmt"
)

func addSymbols(length int, symbol string) string {
	text := ""
	for i := 0; i < length; i++ {
		text += symbol
	}
	return text
}

func main() {
	var length int
	var stair string
	fmt.Scanf("%v", &length)

	for i, j := 1, length-1; i <= length; i++ {
		stair = addSymbols(j, " ") + addSymbols(i, "#")
		fmt.Println(stair)
		j--
	}

}
