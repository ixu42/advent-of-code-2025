package main

import (
    "fmt"
    "os"
	"strings"
	"strconv"
)

func main() {
    data, err := os.ReadFile("inputs/day06.txt")
    if err != nil {
        fmt.Println("Error reading file:", err)
        return
    }

	lines := strings.Split(string(data), "\n")

	// part1
	// num0 := []int{}
	// num1 := []int{}
	// num2 := []int{}
	// num3 := []int{}
	// ops := []string{}
	// for i, line := range lines {
	// 	fields := strings.Fields(line)
	// 	for _, f := range fields {
	// 		if i == 4 {
	// 			ops = append(ops, f)
	// 			continue
	// 		}
	// 		n, err := strconv.Atoi(f)
	// 		if err != nil {
	// 			fmt.Println("Error converting to int:", err)
	// 			return
	// 		}
	// 		if i == 0 {
	// 			num0 = append(num0, n)
	// 			continue
	// 		}
	// 		if i == 1 {
	// 			num1 = append(num1, n)
	// 			continue
	// 		}
	// 		if i == 2 {
	// 			num2 = append(num2, n)
	// 			continue
	// 		}
	// 		num3 = append(num3, n)	
	// 	}
	// }

	// res := 0
	// for i, op := range ops {
	// 	switch op {
	// 		case "+":
	// 			res += num0[i] + num1[i] + num2[i] + num3[i]
	// 		case "*":
	// 			res += num0[i] * num1[i] * num2[i] * num3[i]
	// 	}
	// }

	// fmt.Println("part1:", res)

	// part2
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
		// calculate for prev problem when seeing an operator
		if grid[4][i] != ' ' {
			switch op {
				case '+':
					sum := 0
					for _, n := range nums {
						sum += n
					}
					println("sum:", sum)
					res += sum
				case '*':
					prod := 1
					for _, n := range nums {
						prod *= n
					}
					println("prod:", prod)
					res += prod
				default:
					println("something went wrong")
			}
			nums = []int{}
			numStr = ""
			op = grid[4][i]
		}

		numStr = string(char) + string(grid[1][i]) + string(grid[2][i]) + string(grid[3][i])
		numStr = strings.TrimSpace(numStr)
		if numStr == "" {
			continue
		}
		println("i:", i)
		println("op:", string(op))
		println("numStr:", numStr)
		num, err := strconv.Atoi(numStr)
		if err != nil {
			fmt.Println("Error converting to int:", err)
			return
		}
		println("num:", num)
		nums = append(nums, num)
	}
	switch op {
		case '+':
			sum := 0
			for _, n := range nums {
				sum += n
			}
			println("sum:", sum)
			res += sum
		case '*':
			prod := 1
			for _, n := range nums {
				prod *= n
			}
			println("prod:", prod)
			res += prod
	}
	fmt.Println("part2:", res)
}
