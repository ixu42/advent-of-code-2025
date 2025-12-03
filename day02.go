package main

import (
    "fmt"
    "os"
	"strings"
	"strconv"
)

func isRepeatedTwice(n int) bool {
	str := strconv.Itoa(n)

	if len(str) % 2 != 0 {
		return false
	}

	mid := len(str) / 2
	return str[:mid] == str[mid:]
}

func isRepeatedAtLeastTwice(n int) bool {
	str := strconv.Itoa(n)
	maxRepeat := len(str)

	for i := 2; i <= maxRepeat; i++ {
		if len(str) % i != 0 {
			continue
		}

		chunkSize := len(str) / i
		pattern := str[:chunkSize]
		matched := true

		for j := 1; j < i; j++ {
			if str[j*chunkSize:(j+1)*chunkSize] != pattern {
				matched = false
				break
			}
		}

		if matched {
			return true
		}
	}
	return false
}

func main() {
    data, err := os.ReadFile("inputs/day02.txt")
    if err != nil {
        fmt.Println("Error reading file:", err)
        return
    }

	pairs := strings.Split(string(data), ",")

	// part1
	// res := 0

	// for _, pair := range pairs {
	// 	idPair := strings.Split(pair, "-")
	// 	if len(idPair) != 2 {
	// 		fmt.Println("Invalid id pair:", pair)
	// 		return
	// 	}

	// 	firstId, err1 := strconv.Atoi(idPair[0])
	// 	lastId, err2 := strconv.Atoi(idPair[1])
	// 	if err1 != nil || err2 != nil {
	// 		fmt.Println("Error converting ids:", err1, err2)
	// 		return
	// 	}

	// 	for i := firstId; i <= lastId; i++ {
	// 		if isRepeatedTwice(i) {
	// 			res += i
	// 		}
	// 	}
	// }

	// fmt.Println("part1:", res)

	// part2
	res := 0
	for _, pair := range pairs {
		idPair := strings.Split(pair, "-")
		if len(idPair) != 2 {
			fmt.Println("Invalid id pair:", pair)
			return
		}

		firstId, err1 := strconv.Atoi(idPair[0])
		lastId, err2 := strconv.Atoi(idPair[1])
		if err1 != nil || err2 != nil {
			fmt.Println("Error converting ids:", err1, err2)
			return
		}

		for i := firstId; i <= lastId; i++ {
			if isRepeatedAtLeastTwice(i) {
				res += i
			}
		}
	}
	fmt.Println("part2:", res)
}
