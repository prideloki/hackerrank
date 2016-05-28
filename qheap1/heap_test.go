package main

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func Test01(t *testing.T) {
	inputs := []string{
		"1 16",
		"1 14",
		"3",
		"2 14",
		"3",
		"1 10",
		"1 8",
		"1 7",
		"1 9",
		"1 3",
		"1 4",
		"1 2",
		"1 1",
	}

	//create a heap
	heap := []int{}
	for _, value := range inputs {
		ops := strings.Split(value, " ")
		switch ops[0] {
		case "1":
			//add
			num, err := strconv.Atoi(ops[1])
			if err != nil {
				t.Fatal(err)
			}
			heap = append(heap, num)
			for i := len(heap) - 1; i >= 0; i-- {
				heapify(heap, i)
			}
		case "2":
			//del
			num, err := strconv.Atoi(ops[1])
			if err != nil {
				t.Fatal(err)
			}
			for index, value := range heap {
				if num == value {
					heap[index] = heap[len(heap)-1]
					//remove the last one
					heap = append(heap[:len(heap)-1])
					for i := len(heap) - 1; i >= 0; i-- {
						heapify(heap, i)
					}
					break
				}
			}
		case "3":
			//print
			fmt.Println(heap[0])
		}
	}
}
