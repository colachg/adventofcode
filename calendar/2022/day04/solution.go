package day04

import (
	"strconv"
	"strings"
)

func puzzle(input string) (int, int) {
	var result, overlap int
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		p := strings.Split(line, ",")

		p0 := strings.Split(p[0], "-")
		p1 := strings.Split(p[1], "-")

		a, _ := strconv.Atoi(p0[0])
		b, _ := strconv.Atoi(p0[1])
		c, _ := strconv.Atoi(p1[0])
		d, _ := strconv.Atoi(p1[1])

		if (a <= c && b >= d) || (a >= c && b <= d) {
			result += 1
		}

		if a > d || b < c {
			overlap += 1
		}
	}
	result2 := len(lines) - overlap
	//fmt.Println(result, result2)
	return result, result2

}
