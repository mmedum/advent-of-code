package main

import (
    "fmt"
    "os"
    "bufio"
    "log"
    "strconv"
)


func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {
    var result = 0
    seen_numbers := make(map[int]struct{})
    seen_numbers[result] = struct{}{}
    var seen_twice = false

    for !seen_twice {
        file, err := os.Open("input")
        check(err)
        defer file.Close()

        scanner := bufio.NewScanner(file)
        for scanner.Scan() && !seen_twice {
            input := scanner.Text()
            opr := input[0:1]
            num, err := strconv.Atoi(input[1:len(input)])
            check(err)
            if opr == "-" {
                result -= num
            } else {
                result += num
            }
            if _, ok := seen_numbers[result]; ok {
                fmt.Println("Seen", result)
                seen_twice = true
            } else {
                seen_numbers[result] = struct{}{}
            }
        }
        if err := scanner.Err(); err != nil {
            log.Fatal(err)
        }
    }


    fmt.Println(result)

}
