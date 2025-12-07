package main

import (
    "fmt"
    "os"
	"strings"
)

func main() {
    data, err := os.ReadFile("inputs/day04.txt")
    if err != nil {
        fmt.Println("Error reading file:", err)
        return
    }

	lines := strings.Split(string(data), "\n")

	grid := make([][]rune, len(lines))
	grid_cpy := make([][]rune, len(lines))
	for i, line := range lines {
		grid[i] = []rune(line)
		grid_cpy[i] = []rune(line)
	}

	for _, row := range grid {
		fmt.Println(string(row))
	}

	dirs := [][2]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}

	res := 0
	for r := range grid {
		for c := range grid[r] {
			if grid[r][c] != '@' {
				continue
			}
			count := 0
			for _, d := range dirs {
				nr := r + d[0]
				nc := c + d[1]

				if nr >= 0 && nr < len(grid) && nc >= 0 && nc < len(grid[0]) {
					neighbor := grid[nr][nc]
					if neighbor == '@' {
						count++
					}
				}
			}
			if count < 4 {
				res += 1
				grid_cpy[r][c] = 'x'
			}
		}
	}

	fmt.Println("-----")
	for _, row := range grid_cpy {
		fmt.Println(string(row))
	}

	fmt.Println("part1:", res)
}
