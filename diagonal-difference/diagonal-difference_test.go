package main

import (
	"fmt"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	var capacity, length, left, right, result int
	capacity, length = 3, 3
	rows := make([]string, 0, capacity) // if []string,3 => ["","",""]
	// slice : make([]string,length,capacity)
	// ca

	rows = append(rows, "1 2 3\n")
	rows = append(rows, "1 2 3\n")
	rows = append(rows, "1 2 3\n")
	// rows[0] = "1 2 3\n"
	// rows[1] = "1 2 3\n"
	// rows[2] = "1 2 3\n"
	arg := make([]int, 10, 100) // [0 0 0 0 0 0 0 0 0 0]
	fmt.Println(arg)
	for i, j, k := 0, 0, length-1; i < length; i++ {
		cells := strings.Split(strings.Trim(rows[i], "\n"), " ")
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
	assert.Equal(t, 0, result, "diff = 0 ")

}
