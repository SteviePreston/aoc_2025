package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func solvePass(input []string) int16 {
	curr := int32(50)
	passCount := int16(0)

	for _, i := range input {
		if len(i) < 2 {
			continue
		}

		num, _ := strconv.Atoi(string(i[1:]))

		if strings.ToLower(string(i[0])) == "r" {
			curr = curr + int32(num)
		} else if strings.ToLower(string(i[0])) == "l" {
			curr = curr - int32(num)
		} else {
			fmt.Println("No Valid Input")
			continue
		}

		curr = ((curr % 100) + 100) % 100
		if curr == 0 {
			passCount += 1
		}
	}
	return passCount
}

func main() {
	inputFile, _ := os.ReadFile("input.txt")
	input := strings.Split(string(inputFile), "\n")

	ans := solvePass(input)
	fmt.Printf("The password is: %d\n", ans)

}
