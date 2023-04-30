package day12

import (
	"strings"
)

type Point struct {
	line int
	col  int
}

func part1(input string) int {

	tmp := strings.Split(strings.TrimSpace(input), "\n")
	puzzle := make([][]byte, len(tmp))

	// table to maintenance visited or not
	visited := make([][]int, len(tmp))

	// init the visited table and parse "S" and "E"
	// parse "S" and "E"
	var start, end Point
	for line := range tmp {
		visited[line] = make([]int, len(tmp[line]))
		puzzle[line] = make([]byte, len(tmp[line]))

		for col := range tmp[line] {
			puzzle[line][col] = tmp[line][col]
			if tmp[line][col] == 'S' {
				start = Point{line, col}
				visited[line][col] = 1
				puzzle[line][col] = byte('a')
			}
			if tmp[line][col] == 'E' {
				end = Point{line, col}
				puzzle[line][col] = byte('z')
			}
		}
	}

	// visit from the "S" point
	queue := []Point{start}
	for {
		p := queue[0]

		// current height
		height, l, c := puzzle[p.line][p.col], p.line, p.col

		// check if it's end point
		if p == end {
			//for _, v := range visited {
			//	fmt.Println(v)
			//}
			return visited[l][c] - 1
		}

		// left and not visited
		if c-1 >= 0 && visited[l][c-1] == 0 && height+1 >= puzzle[l][c-1] {
			visited[l][c-1] = visited[l][c] + 1
			queue = append(queue, Point{l, c - 1})
		}

		// up and not visited
		if l-1 >= 0 && visited[l-1][c] == 0 && height+1 >= puzzle[l-1][c] {
			visited[l-1][c] = visited[l][c] + 1
			queue = append(queue, Point{l - 1, c})
		}

		// right and not visited
		if c+1 < len(puzzle[l]) && visited[l][c+1] == 0 && height+1 >= puzzle[l][c+1] {
			visited[l][c+1] = visited[l][c] + 1
			queue = append(queue, Point{l, c + 1})
		}

		// down and not visited
		if l+1 < len(puzzle) && visited[l+1][c] == 0 && height+1 >= puzzle[l+1][c] {
			visited[l+1][c] = visited[l][c] + 1
			queue = append(queue, Point{l + 1, c})
		}
		queue = queue[1:]
	}
}

func part2(input string) int {
	tmp := strings.Split(strings.TrimSpace(input), "\n")
	puzzle := make([][]byte, len(tmp))

	// table to maintenance visited or not
	visited := make([][]int, len(tmp))
	// init the visited table and parse "S" and "E"
	// parse "S" and "E"
	var end Point
	for line := range tmp {
		visited[line] = make([]int, len(tmp[line]))
		puzzle[line] = make([]byte, len(tmp[line]))

		for col := range tmp[line] {
			puzzle[line][col] = tmp[line][col]
			if tmp[line][col] == 'S' {
				puzzle[line][col] = byte('a')
			}
			if tmp[line][col] == 'E' {
				end = Point{line, col}
				visited[line][col] = 1
				puzzle[line][col] = byte('z')
			}
		}
	}

	// visit from the "S" point
	queue := []Point{end}
	for {
		p := queue[0]

		// current height
		height, l, c := puzzle[p.line][p.col], p.line, p.col

		// check if it's end point
		if 'a' == puzzle[l][c] {
			//for _, v := range visited {
			//	fmt.Println(v)
			//}
			return visited[l][c] - 1
		}

		// left and not visited
		if c-1 >= 0 && visited[l][c-1] == 0 && height-1 <= puzzle[l][c-1] {
			visited[l][c-1] = visited[l][c] + 1
			queue = append(queue, Point{l, c - 1})
		}

		// up and not visited
		if l-1 >= 0 && visited[l-1][c] == 0 && height-1 <= puzzle[l-1][c] {
			visited[l-1][c] = visited[l][c] + 1
			queue = append(queue, Point{l - 1, c})
		}

		// right and not visited
		if c+1 < len(puzzle[l]) && visited[l][c+1] == 0 && height-1 <= puzzle[l][c+1] {
			visited[l][c+1] = visited[l][c] + 1
			queue = append(queue, Point{l, c + 1})
		}

		// down and not visited
		if l+1 < len(puzzle) && visited[l+1][c] == 0 && height-1 <= puzzle[l+1][c] {
			visited[l+1][c] = visited[l][c] + 1
			queue = append(queue, Point{l + 1, c})
		}
		queue = queue[1:]
	}
}
