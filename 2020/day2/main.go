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
		pos1, _ := strconv.Atoi(string(rangeElement[0]))
		pos2, _ := strconv.Atoi(string(rangeElement[1]))
		char := string(element[1][0])
		password := string(element[2])

		if string(password[pos1-1]) == char && string(password[pos2-1]) != char {
			valid++
		} else if string(password[pos1-1]) != char && string(password[pos2-1]) == char {
			valid++
		}
	}

	fmt.Println(valid)
}
