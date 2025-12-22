package main

import (
    "fmt"
    "os"
	"strings"
	"log"
)

var graph map[string][]string

func countPath(current string) int {
	if current == "out" {
		return 1
	}

	totalPaths := 0
	for _, next := range graph[current] {
		totalPaths += countPath(next)
	}
	return totalPaths
}

func main() {
    data, err := os.ReadFile("inputs/day11.txt")
    if err != nil {
        fmt.Println("Error reading file:", err)
        return
    }

	lines := strings.Split(string(data), "\n")

	graph = make(map[string][]string)
	for _, line := range lines {
		parts := strings.Split(line, ":")
		if len(parts) != 2 {
			log.Fatalf("Invalid line format: %s", line)
		}
		device := strings.TrimSpace(parts[0])
		outputs := strings.TrimSpace(parts[1])
		outputList := strings.Fields(outputs)
		graph[device] = outputList
	}

	// for device, outs := range graph {
	// 	fmt.Println(device, "->", outs)
	// }

	fmt.Println("part1:", countPath("you"))
}
