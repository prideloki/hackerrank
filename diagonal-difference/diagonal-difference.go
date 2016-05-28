package main

// how to  write test for stdin
import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var length, left, right, result int

	fmt.Scanln(&length)
	//input length
	//input 3 lines
	//save to 3 lines then calculate
	in := bufio.NewReader(os.Stdin)
	for i, j, k := 0, 0, length-1; i < length; i++ {
		line, _ := in.ReadString('\n')
		cells := strings.Split(strings.Trim(line, "\n"), " ")
		leftP, _ := strconv.Atoi(cells[j])
		rightP, _ := strconv.Atoi(cells[k])
		left += leftP
		right += rightP

		j++
		k--
	}
	result = left - right
	if result < 0 {
		result = -result
	}
	fmt.Println(result)
}
