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

func up(heap []int, j int) {
	for {
		i := (j - 1) / 2 //parent node
		// if the new pos is already less than
		if i == j || !(heap[j] < heap[i]) {
			break
		}
		heap[i], heap[j] = heap[j], heap[i]
		j = i
	}
}

//down which child to go, left or right
func down(heap []int, i, n int) {
	for {
		j1 := 2*i + 1 // left child
		// shouldn't exceed the n and less than 0
		if j1 >= n || j1 < 0 {
			break
		}
		j := j1 // left child
		// j2 is right child(j1+1)
		// j will be j2 if h.j1 > h.j2
		if j2 := j1 + 1; j2 < n && !(heap[j1] < heap[j2]) {
			j = j2 // = 2*i + 2  // right child
		}
		// if j is less than then it's done
		if !(heap[j] < heap[i]) {
			break
		}
		heap[i], heap[j] = heap[j], heap[i]
		i = j
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
			//try price for each step
			min = i
		} else {
			break
		}
	}
}
