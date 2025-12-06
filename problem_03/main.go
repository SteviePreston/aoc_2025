package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"strings"
)

func maxJoltagePart1(bank string) int {
	maxJolt := 0

	for i := 0; i < len(bank); i++ {
		for j := i + 1; j < len(bank); j++ {
			jolt := int(bank[i]-'0')*10 + int(bank[j]-'0')
			if jolt > maxJolt {
				maxJolt = jolt
			}
		}
	}

	return maxJolt
}

func maxJoltagePart2(bank string, numBatteries int) string {
	selected := make([]rune, 0, numBatteries)
	pos := 0

	for len(selected) < numBatteries {
		needed := numBatteries - len(selected)
		remaining := len(bank) - pos
		maxSearchPos := pos + remaining - needed
		maxIdx := pos

		for i := pos; i <= maxSearchPos; i++ {
			if bank[i] > bank[maxIdx] {
				maxIdx = i
			}
		}

		selected = append(selected, rune(bank[maxIdx]))
		pos = maxIdx + 1
	}
	return string(selected)
}

func mainp1() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var totalJoltage int = 0

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		jolt := maxJoltagePart1(line)
		fmt.Printf("%s -> %d\n", line, jolt)

		totalJoltage += jolt
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("\nTotal output joltage: %d\n", totalJoltage)
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	totalJoltage := big.NewInt(0)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		jolt := maxJoltagePart2(line, 12)
		fmt.Printf("%s -> %s\n", line, jolt)

		joltBig := big.NewInt(0)
		joltBig.SetString(jolt, 10)
		totalJoltage.Add(totalJoltage, joltBig)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("\nTotal output joltage: %s\n", totalJoltage.String())
}
