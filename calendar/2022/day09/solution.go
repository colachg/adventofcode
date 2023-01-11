package main

import (
	"strconv"
	"strings"
)

type Point struct{ X, Y int }

func (p Point) Follow(h Point) Point {
	diffX, diffY := h.X-p.X, h.Y-p.Y
	if abs(diffX) <= 1 && abs(diffY) <= 1 {
		return p
	}
	return Point{p.X + sign(diffX), p.Y + sign(diffY)}
}

func abs(i int) int {
	if i >= 0 {
		return i
	} else {
		return -i
	}
}

func sign(i int) int {
	if i == 0 {
		return 0
	}
	if i > 0 {
		return 1
	}
	return -1
}

func part1(input string) int {
	var head, tail Point
	visited := make(map[Point]int)

	for _, line := range strings.Split(input, "\n") {
		tmp := strings.Split(line, " ")
		step, _ := strconv.Atoi(tmp[1])
		for i := 0; i < step; i++ {
			switch tmp[0] {
			case "U":
				head.Y++
			case "D":
				head.Y--
			case "L":
				head.X--
			case "R":
				head.X++
			}
			tail = tail.Follow(head)
			visited[tail] += 1
		}
	}
	return len(visited)
}

func part2(input string) int {
	rope := make([]Point, 10)
	visited := make(map[Point]int)

	for _, line := range strings.Split(input, "\n") {
		tmp := strings.Split(line, " ")
		step, _ := strconv.Atoi(tmp[1])
		for i := 0; i < step; i++ {
			switch tmp[0] {
			case "U":
				rope[0].Y++
			case "D":
				rope[0].Y--
			case "L":
				rope[0].X--
			case "R":
				rope[0].X++
			}

			for j := 1; j < len(rope); j++ {
				rope[j] = rope[j].Follow(rope[j-1])
			}

			visited[rope[9]] += 1
		}
	}
	return len(visited)
}
