package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	f, err := os.Open("input")
	check(err)
	defer f.Close()

	scanner := bufio.NewScanner(f)

	totalFuelRequirement := 0
	for scanner.Scan() {
		moduleMass, err := strconv.Atoi(scanner.Text())
		check(err)

		fuel := (moduleMass / 3) - 2
		totalFuelRequirement += fuel
	}

	fmt.Println("Total fuel reuqirement:", totalFuelRequirement)
}
