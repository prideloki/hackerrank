package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	var length int
	fmt.Scanf("%v", &length)

	in := bufio.NewReader(os.Stdin)
	heap := []int{}
	for inputLine := 0; inputLine < length; inputLine++ {
		line, _ := in.ReadString('\n')
		ops := strings.Split(strings.Trim(line, "\n"), " ")
		switch ops[0] {
		case "1":
			//add
			num, err := strconv.Atoi(ops[1])
			if err != nil {
				log.Fatal(err)
			}
			heap = append(heap, num)
			for i := len(heap) - 1; i >= 0; i-- {
				heapify(heap, i)
			}

		case "2":
			//del
			num, err := strconv.Atoi(ops[1])
			if err != nil {
				log.Fatal(err)
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
			if len(heap) > 0 {
				fmt.Println(heap[0])
			} else {
				fmt.Print(0)
			}

		}
	}
}

func heapify(heap []int, i int) {
	heapSize := len(heap) - 1
	min := i
	for {
		l := 2 * i
		r := 2*i + 1
		if l <= heapSize && heap[l] < heap[i] {
			min = l
		}
		if r <= heapSize && heap[r] < heap[i] {
			min = r
		}
		if min != i {
			heap[i], heap[min] = heap[min], heap[i]
			min = i
		} else {
			break
		}
	}
}
