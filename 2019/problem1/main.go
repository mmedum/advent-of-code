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

		totalFuelRequirement += calculateFuel(moduleMass)
	}

	fmt.Println("Total fuel requirement:", totalFuelRequirement)
}

func calculateFuel(mass int) int {
	subMass := (mass / 3) - 2
	if subMass <= 0 {
		return 0
	}
	return subMass + calculateFuel(subMass)
}
