package main

import (
    "fmt"
    "os"
	"strings"
)

func getMaxJoltage (s string) int {
	digits := []byte(s)
	maxJoltage := -1

	for i := 0; i < len(digits) - 1; i++ {
		d1 := int(digits[i] - '0')
		for j := i + 1; j < len(digits); j++ {
			d2 := int(digits[j] - '0')
			val := d1 * 10 + d2
			if val > maxJoltage {
				maxJoltage = val
			}
		}
	}

	return maxJoltage
}

func main() {
    data, err := os.ReadFile("inputs/day03.txt")
    if err != nil {
        fmt.Println("Error reading file:", err)
        return
    }

	banks := strings.Split(string(data), "\n")

	// part1
	res := 0
	for _, bank := range banks {
		res += getMaxJoltage(bank)
	}

	fmt.Println("part1:", res)
}
