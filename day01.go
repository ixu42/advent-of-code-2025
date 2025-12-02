package main

import (
    "fmt"
    "os"
	"strings"
	"strconv"
)

func parseRotation(line string) (direction byte, distance int, err error) {
	direction = line[0]
	distance, err = strconv.Atoi(line[1:])
	return direction, distance, err
}

func rotateAndCount(pos int, direction byte, distance int) (int, int) {
	countAtZero := 0
	for i := 0; i < distance; i++ {
		if direction == 'L' {
			pos--
			if pos < 0 {
				pos = 99
			}
		} else if direction == 'R' {
			pos++
			if pos > 99 {
				pos = 0
			}
		}
		if pos == 0 {
			countAtZero++
		}
	}
	return pos, countAtZero
}

func solvePart1(lines []string) int {
	pos := 50
	countAtZero := 0

	for _, line := range lines {
		if line == "" {
			continue
		}

		direction, distance, err := parseRotation(line)
		if err != nil {
			fmt.Println("Error parsing rotation:", err)
			return 0
		}

		pos, _ = rotateAndCount(pos, direction, distance)

		if pos == 0 {
			countAtZero++
		}
	}
	return countAtZero
}

func solvePart2(lines []string) int {
	pos := 50
	countAtZero := 0

	for _, line := range lines {
		if line == "" {
			continue
		}

		direction, distance, err := parseRotation(line)
		if err != nil {
			fmt.Println("Error parsing rotation:", err)
			return 0
		}

		tempCount := 0
		pos, tempCount = rotateAndCount(pos, direction, distance)
		countAtZero += tempCount
	}
	return countAtZero
}

func main() {
    data, err := os.ReadFile("inputs/day01.txt")
    if err != nil {
        fmt.Println("Error reading file:", err)
        return
    }

	lines := strings.Split(string(data), "\n")

	fmt.Println("part1:", solvePart1(lines))
	fmt.Println("part2:", solvePart2(lines))
}
