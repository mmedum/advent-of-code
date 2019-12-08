package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// Opcodes
// 1 = addition, pos1, pos2, storePos
// 2 = multiply, pos1, pos2, storePost
// 99 = halt
// After 1 || 2 move 4 positions forward
// Output position 0, after replacing
// position 1 with 12 and position 2 with 2
func main() {
	f, err := os.Open("input")
	check(err)
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		memory := strings.Split(scanner.Text(), ",")

		for noun := 0; noun <= 99; noun++ {
			for verb := 0; verb <= 99; verb++ {
				memoryCopy := make([]string, len(memory))
				copy(memoryCopy, memory)
				memoryCopy[1] = strconv.Itoa(noun)
				memoryCopy[2] = strconv.Itoa(verb)
				result := compute(memoryCopy)
				if result == 19690720 {
					fmt.Println(noun, verb)
					fmt.Println(100*noun + verb)
					os.Exit(0)
				}
			}
		}
	}
}

func compute(memory []string) int {
	for i := 0; i < len(memory); i++ {
		operation, _ := strconv.Atoi(memory[i])

		if operation == 99 {
			result, _ := strconv.Atoi(memory[0])
			return result
		}

		address1, _ := strconv.Atoi(memory[i+1])
		address2, _ := strconv.Atoi(memory[i+2])
		storeAddress, _ := strconv.Atoi(memory[i+3])
		value1, _ := strconv.Atoi(memory[address1])
		value2, _ := strconv.Atoi(memory[address2])
		if operation == 1 {
			memory[storeAddress] = strconv.Itoa(value1 + value2)
		} else {
			memory[storeAddress] = strconv.Itoa(value1 * value2)
		}
		i += 3
	}
	return 0
}
