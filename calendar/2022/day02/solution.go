package day02

import (
	"bytes"
)

func part1(input []byte) int {
	rmap := map[string]int{
		"A X": 3 + 1,
		"A Y": 6 + 2,
		"A Z": 0 + 3,
		"B X": 0 + 1,
		"B Y": 3 + 2,
		"B Z": 6 + 3,
		"C X": 6 + 1,
		"C Y": 0 + 2,
		"C Z": 3 + 3,
	}
	var sum int
	lines := bytes.Split(input, []byte{'\n'})
	for _, line := range lines {
		sum += rmap[string(line)]
	}
	return sum
}

func part2(input []byte) int {
	rmap := map[string]int{
		"A X": 0 + 3,
		"A Y": 3 + 1,
		"A Z": 6 + 2,
		"B X": 0 + 1,
		"B Y": 3 + 2,
		"B Z": 6 + 3,
		"C X": 0 + 2,
		"C Y": 3 + 3,
		"C Z": 6 + 1,
	}

	var sum int
	lines := bytes.Split(input, []byte{'\n'})
	for _, line := range lines {
		sum += rmap[string(line)]
	}
	return sum
}
