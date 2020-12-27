package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	fileBytes, _ := ioutil.ReadFile("input")

	data := strings.Split(string(fileBytes), "\n")
	valid := 0

	for i := 0; i < len(data)-1; i++ {
		element := strings.Split(data[i], " ")
		rangeElement := strings.Split(string(element[0]), "-")
		min, _ := strconv.Atoi(string(rangeElement[0]))
		max, _ := strconv.Atoi(string(rangeElement[1]))
		char := string(element[1][0])
		password := string(element[2])

		count := 0
		for _, c := range password {
			if string(c) == char {
				count++
			}
		}

		if count >= min && count <= max {
			valid++
		}
	}

	fmt.Println(valid)
}
