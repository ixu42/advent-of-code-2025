package main

import (
    "fmt"
    "os"
	"strings"
	"strconv"
)

func solvePart1(lines []string) int {
	num0 := []int{}
	num1 := []int{}
	num2 := []int{}
	num3 := []int{}
	ops := []string{}
	for i, line := range lines {
		fields := strings.Fields(line)
		for _, f := range fields {
			if i == 4 {
				ops = append(ops, f)
				continue
			}
			n, err := strconv.Atoi(f)
			if err != nil {
				fmt.Println("Error converting to int:", err)
				return 0
			}
			if i == 0 {
				num0 = append(num0, n)
				continue
			}
			if i == 1 {
				num1 = append(num1, n)
				continue
			}
			if i == 2 {
				num2 = append(num2, n)
				continue
			}
			num3 = append(num3, n)	
		}
	}

	res := 0
	for i, op := range ops {
		switch op {
			case "+":
				res += num0[i] + num1[i] + num2[i] + num3[i]
			case "*":
				res += num0[i] * num1[i] * num2[i] * num3[i]
		}
	}
	return res
}

func calculate(nums []int, op string) int {
	res := 0
	switch op {
		case "+":
			sum := 0
			for _, n := range nums {
				sum += n
			}
			res += sum
		case "*":
			prod := 1
			for _, n := range nums {
				prod *= n
			}
			res += prod
	}
	return res
}

func solvePart2(lines []string) int {
	res := 0
	grid := make([][]rune, len(lines))
	for i, line := range lines {
		grid[i] = make([]rune, len(line))
		grid[i] = []rune(line)
	}

	nums := []int{}
	numStr := ""
	op := ' '
	for i := 0; i < len(grid[0]); i++ {
		char := grid[0][i]
		// calculate for prev problem when seeing a new operator
		NewOp := grid[4][i]
		if NewOp != ' ' {
			res += calculate(nums, string(op))
			nums = []int{}
			numStr = ""
			op = NewOp
		}

		numStr = string(char) + string(grid[1][i]) + string(grid[2][i]) + string(grid[3][i])
		numStr = strings.TrimSpace(numStr)
		if numStr == "" {
			continue
		}
		num, err := strconv.Atoi(numStr)
		if err != nil {
			fmt.Println("Error converting to int:", err)
			return 0
		}
		nums = append(nums, num)
	}
	res += calculate(nums, string(op))
	return res
}

func main() {
    data, err := os.ReadFile("inputs/day06.txt")
    if err != nil {
        fmt.Println("Error reading file:", err)
        return
    }

	lines := strings.Split(string(data), "\n")

	fmt.Println("part1:", solvePart1(lines))
	fmt.Println("part2:", solvePart2(lines))
}
