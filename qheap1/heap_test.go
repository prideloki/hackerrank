package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"testing"
)

func Test01(t *testing.T) {
	inputs := []string{
		"1 486134735",
		"1 551386163",
		"1 681080115",
		"1 98708774",
		"1 676735757",
		"3",
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
				log.Fatal(err)
			}
			heap = append(heap, num)
			up(heap, len(heap)-1)

		case "2":
			//del
			num, err := strconv.Atoi(ops[1])
			if err != nil {
				log.Fatal(err)
			}
			for index, value := range heap {
				if num == value {
					n := len(heap) - 1
					heap[index], heap[n] = heap[n], heap[index]
					down(heap, index, n)
					up(heap, index)
					//remove the last one
					heap = append(heap[:len(heap)-1])
					break
				}
			}
		case "3":
			//print
			if len(heap) > 0 {
				fmt.Println(heap[0])
			} else {
				fmt.Print(0)
			}

		}
	}
}
