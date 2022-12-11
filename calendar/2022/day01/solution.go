package day01

import (
	"bytes"
	_ "embed"
	"fmt"
	"sort"
	"strconv"
)

func part1(input []byte) int {
	var i, j int
	words := bytes.Split(input, []byte{'\n'})
	for _, w := range words {
		if bytes.Equal(w, []byte{}) {
			if i > j {
				j = i
			}
			i = 0
		} else {
			num, _ := strconv.Atoi(string(w))
			i += num
		}
	}
	return j
}

func part2(input []byte) {
	var i int
	var sums []int
	words := bytes.Split(input, []byte{'\n'})
	for _, w := range words {
		if bytes.Equal(w, []byte{}) {
			sums = append(sums, i)
			i = 0
		} else {
			num, _ := strconv.Atoi(string(w))
			i += num
		}
	}
	sort.Ints(sums)
	fmt.Println(sums[len(sums)-1] + sums[len(sums)-2] + sums[len(sums)-3])
}
