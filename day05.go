package main

import (
    "fmt"
    "os"
	"strings"
	"strconv"
)

type Range struct {
	start int
	end   int
}

func isFresh(id int, ranges []Range) bool {
	for _, r := range ranges {
		if id >= r.start && id <= r.end {
			return true
		}
	}
	return false
}

func main() {
    input, err := os.ReadFile("inputs/day05.txt")
    if err != nil {
        fmt.Println("Error reading file:", err)
        return
    }

	var ranges []Range
	var ids []int

	lines := strings.Split(string(input), "\n")

	for _, line := range lines {
		if line == "" {
			continue
		}

		if strings.Contains(line, "-") {
			parts := strings.Split(line, "-")
			start, err1 := strconv.Atoi(parts[0])
			end, err2 := strconv.Atoi(parts[1])
			if err1 != nil || err2 != nil {
				fmt.Println("Error converting to integer:", err1, err2)
				return
			}
			ranges = append(ranges, Range{start: start, end: end})
		} else {
			val, err := strconv.Atoi(line)
			if err != nil {
				fmt.Println("Error converting to integer:", err)
				return
			}
			ids = append(ids, val)
		}
	}
	// fmt.Println("Ranges:", ranges)
	// fmt.Println("Ints:", ids)

	// part1
	res := 0
	for _, id := range ids {
		if isFresh(id, ranges) {
			res++
		}
	}
	fmt.Println("part1:", res)
}
