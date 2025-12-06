package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func isInvalidIDPart1(id string) bool {
	if len(id)%2 != 0 {
		return false
	}

	half := len(id) / 2
	firstHalf := id[:half]
	secondHalf := id[half:]

	return firstHalf == secondHalf
}

func isInvalidIDPart2(id string) bool {
	for patternLen := 1; patternLen <= len(id)/2; patternLen++ {
		if len(id)%patternLen != 0 {
			continue
		}

		pattern := id[:patternLen]
		valid := true

		for i := 0; i < len(id); i += patternLen {
			if id[i:i+patternLen] != pattern {
				valid = false
				break
			}
		}

		if valid {
			return true
		}
	}

	return false
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var totalSum int64 = 0

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		ranges := strings.Split(line, ",")

		for _, r := range ranges {
			parts := strings.Split(strings.TrimSpace(r), "-")
			if len(parts) != 2 {
				continue
			}

			start, err := strconv.ParseInt(parts[0], 10, 64)
			if err != nil {
				continue
			}

			end, err := strconv.ParseInt(parts[1], 10, 64)
			if err != nil {
				continue
			}

			for id := start; id <= end; id++ {
				idStr := strconv.FormatInt(id, 10)
				if isInvalidIDPart2(idStr) {
					fmt.Printf("%d is invalid\n", id)
					totalSum += id
				}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("\nTotal sum: %d\n", totalSum)
}
