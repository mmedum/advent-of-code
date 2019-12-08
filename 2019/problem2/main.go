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
		commands := strings.Split(scanner.Text(), ",")
		fmt.Println(commands)

		for i := 0; i < len(commands); i++ {
			command, _ := strconv.Atoi(commands[i])

			if command == 99 {
				break
			}

			postion1, _ := strconv.Atoi(commands[i+1])
			postion2, _ := strconv.Atoi(commands[i+2])
			storePostion, _ := strconv.Atoi(commands[i+3])
			value1, _ := strconv.Atoi(commands[postion1])
			value2, _ := strconv.Atoi(commands[postion2])
			if command == 1 {
				commands[storePostion] = strconv.Itoa(value1 + value2)
			} else {
				commands[storePostion] = strconv.Itoa(value1 * value2)
			}
			i += 3
		}
		fmt.Println(commands[0])
	}
}
