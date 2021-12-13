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

func partA(sliceData []string) string {
	for i := 0; i < len(sliceData)-2; i++ {
		for j := i + 1; j < len(sliceData)-1; j++ {
			a, _ := strconv.Atoi(sliceData[i])
			b, _ := strconv.Atoi(sliceData[j])
			if a+b == 2020 {
				return strconv.Itoa(a * b)
			}
		}
	}
	return "Not Found"
}

func partB(sliceData []string) string {
	for i := 0; i < len(sliceData)-2; i++ {
		for j := i + 1; j < len(sliceData)-1; j++ {
			for x := j + 1; x < len(sliceData); x++ {
				a, _ := strconv.Atoi(sliceData[i])
				b, _ := strconv.Atoi(sliceData[j])
				c, _ := strconv.Atoi(sliceData[x])
				if a+b+c == 2020 {
					return strconv.Itoa(a * b * c)
				}
			}
		}
	}
	return "Not Found"
}

func main() {
	fileBytes, err := ioutil.ReadFile("../input")
	check(err)

	sliceData := strings.Split(string(fileBytes), "\n")

	a := partA(sliceData)
	b := partB(sliceData)

	fmt.Println("A: " + a + ", B: " + b)
}
