package main

import (
    "fmt"
    "os"
	"strings"
	"strconv"
)

type pos struct {
	r, c int
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}

func main() {
    data, err := os.ReadFile("inputs/day09.txt")
    if err != nil {
        fmt.Println("Error reading file:", err)
        return
    }

	lines := strings.Split(string(data), "\n")

	redTiles := make([]pos, 0, len(lines))
	for _, line := range lines {
		elems := strings.Split(line, ",")
		if len(elems) != 2 {
			fmt.Println("invalid position:", line)
			continue
		}
		c, errC := strconv.Atoi(elems[0])
		r, errR := strconv.Atoi(elems[1])
		if errC != nil || errR != nil {
			fmt.Println("invalid coordinates:", line)
			continue
		}
		redTiles = append(redTiles, pos{r, c})
	}

	maxSize := 0
	for i := 0; i < len(redTiles); i++ {
		for j := i + 1; j < len(redTiles); j++ {
			size := (abs(redTiles[i].r - redTiles[j].r) + 1) * (abs(redTiles[i].c - redTiles[j].c) + 1)
			if size > maxSize {
				maxSize = size
			}
		}
	}

	fmt.Printf("part1: %d\n", maxSize)

}
