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
				return
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

	fmt.Println("part1:", res)
}
