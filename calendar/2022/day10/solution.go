package day10

import (
	"fmt"
	"strings"
)

func parse(input string) []int {
	X := 1
	result := []int{1}

	for _, line := range strings.Split(strings.TrimSpace(input), "\n") {
		var opt string
		var num int

		result = append(result, X)
		if _, err := fmt.Sscanf(line, "%s %d", &opt, &num); err == nil {
			X += num
			result = append(result, X)
		}
	}
	return result
}

func part1(input string) int {
	result := parse(input)
	var sum int
	for i := 19; i <= len(result); i += 40 {
		sum += (i + 1) * result[i]
	}
	return sum
}

func part2(input string) int {
	result := parse(input)
	for i := 0; i < 6; i++ {
		for j := 0; j < 40; j++ {
			p := i*40 + j
			if j == result[p] || j-1 == result[p] || j+1 == result[p] {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
	return 0
}
