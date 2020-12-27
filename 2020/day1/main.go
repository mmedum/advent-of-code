package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	fileBytes, err := ioutil.ReadFile("input")
	check(err)

	sliceData := strings.Split(string(fileBytes), "\n")

	for i := 0; i < len(sliceData)-1; i++ {
		for j := i + 1; j < len(sliceData); j++ {
			a, _ := strconv.Atoi(sliceData[i])
			b, _ := strconv.Atoi(sliceData[j])
			if a+b == 2020 {
				fmt.Println(a * b)
				return
			}
		}
	}
}
