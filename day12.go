package main

import (
    "fmt"
    "os"
	"strings"
	"log"
	"bufio"
	"strconv"
)

type shape [][]bool // true = '#', false = '.'

type region struct {
	width, height	int
	counts			[]int // number of presents in each shape
}

var shapes []shape
var regions []region
var tileCountForEach []int // tile count for each shape

func printShapes() {
	for i, shape := range shapes {
		fmt.Printf("Shape %d:\n", i)
		for _, row := range shape {
			for _, cell := range row {
				if cell {
					fmt.Print("#")
				} else {
					fmt.Print(".")
				}
			}
			fmt.Println()
		}
	}
}

func printRegions() {
	for i, region := range regions {
		fmt.Printf("Region %d: %dx%d Counts: %v\n", i, region.width, region.height, region.counts)
	}
}

func countTilesForEachShape() {
	tileCountForEach = make([]int, len(shapes))
	for i, shape := range shapes {
		count := 0
		for _, row := range shape {
			for _, cell := range row {
				if cell {
					count++
				}
			}
		}
		tileCountForEach[i] = count
	}
}

func printTileCountForEach() {
	for i, count := range tileCountForEach {
		fmt.Printf("Shape %d has %d tiles\n", i, count)
	}
}

func parseShapes(line string, currentShape *shape, shapesParsed *bool) {
	if strings.HasSuffix(line, ":") {
		*currentShape = shape{}
	} else if strings.Contains(line, "#") || strings.Contains(line, ".") {
		row := make([]bool, len(line))
		for i, c := range line {
			if c == '#' {
				row[i] = true
			} else {
				row[i] = false
			}
		}
		*currentShape = append(*currentShape, row)
	} else {
		if strings.Contains(line, "x") {
			*shapesParsed = true
		} else {
			shapes = append(shapes, *currentShape)
		}
	}
}

func parseRegions(line string) {
	parts := strings.Split(line, ":")
	dimsStr := strings.TrimSpace(parts[0])
	CountsStr := strings.TrimSpace(parts[1])

	dimParts := strings.Split(dimsStr, "x")
	width, err1 := strconv.Atoi(strings.TrimSpace(dimParts[0]))
	height, err2 := strconv.Atoi(strings.TrimSpace(dimParts[1]))
	if err1 != nil || err2 != nil {
		log.Fatal("invalid region dimensions:", dimsStr)
	}

	countParts := strings.Split(CountsStr, " ")
	counts := make([]int, len(countParts))

	for i, countStr := range countParts {
		count, err := strconv.Atoi(strings.TrimSpace(countStr))
		if err != nil {
			log.Fatal("invalid count:", countStr)
		}
		counts[i] = count
	}

	regions = append(regions, region{
		width:  width,
		height: height,
		counts: counts,
	})
}

func parseInput(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	shapes = make([]shape, 0)
    regions = make([]region, 0)
	shapesParsed := false
	var currentShape shape

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if !shapesParsed {
			parseShapes(line, &currentShape, &shapesParsed)
		}

		if shapesParsed {
			parseRegions(line)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal("input error:", err)
	}

	// printShapes()
	// printRegions()
}

func main() {
	parseInput("inputs/day12.txt")

	countTilesForEachShape()
	// printTileCountForEach()

	nonFitCount := 0
	fitCount := 0
	for _, region := range regions {
		area := region.width * region.height
		tilesNeeded := 0
		maxNeeded := 0
		for j, count := range region.counts {
			tilesNeeded += count * tileCountForEach[j]
			maxNeeded += 3 * 3 * count
		}
		if area < tilesNeeded {
			nonFitCount++
		} else {
			if area >= maxNeeded {
				fitCount++
			} else {
				fmt.Println(" Might fit.")
			}
		}
	}

	// fmt.Printf("total: %d\n", len(regions))
	// fmt.Printf("definitely won't fit: %d\n", nonFitCount)
	// fmt.Printf("definitely will fit: %d\n", fitCount)
	if len(regions) != nonFitCount + fitCount {
		fmt.Println("uncertain cases exist - more complex logic needed to solve this problem")
	} else {
		fmt.Println("result:", fitCount)
	}
}
