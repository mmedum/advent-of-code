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

	for i := 0; i < len(sliceData)-2; i++ {
		for j := i + 1; j < len(sliceData)-1; j++ {
			for x := j + 1; x < len(sliceData); x++ {
				a, _ := strconv.Atoi(sliceData[i])
				b, _ := strconv.Atoi(sliceData[j])
				c, _ := strconv.Atoi(sliceData[x])
				if a+b+c == 2020 {
					fmt.Println(a * b * c)
					return
				}
			}
		}
	}
}
