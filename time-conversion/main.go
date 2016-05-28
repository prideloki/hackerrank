package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var input, time, hourText, timeUnit string
	var hour int
	fmt.Scanf("%s", &input)
	s := ":"
	timeUnit = input[len(input)-2:]
	time = input[:len(input)-2]
	parts := strings.Split(time, s)
	hour, _ = strconv.Atoi(parts[0])
	if timeUnit == "PM" && hour < 12 {
		hour += 12
	}
	if hour == 12 {
		if timeUnit == "AM" {
			hour = 0
		} else if timeUnit == "PM" {
			hour = 12
		}
	}

	if hour < 10 {
		hourText = "0" + strconv.Itoa(hour)
	} else {
		hourText = strconv.Itoa(hour)
	}
	fmt.Printf("%s:%s:%s\n", hourText, parts[1], parts[2])

}
