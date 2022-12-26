package day08

import (
	"fmt"
	"strings"
)

func part1(input string) int {
	// parse input
	lines := strings.Split(input, "\n")
	rows, cols := len(lines), len(lines[0])

	// setup grid
	grids := make([][]bool, rows)
	for i := range grids {
		grids[i] = make([]bool, cols)
	}

	// the edge trees
	result := rows*2 + (cols-2)*2

	for i := 1; i < rows-1; i++ {
		left := lines[i][0]
		right := lines[i][cols-1]

		//from left to right
		for j := 1; j < cols-1; j++ {
			visible := grids[i][j]
			if int(lines[i][j]) > int(left) {
				if !visible {
					result++
					grids[i][j] = true
				}
				left = lines[i][j]
			}
		}

		//from right to left
		for j := cols - 2; j > 0; j-- {
			visible := grids[i][j]
			if int(lines[i][j]) > int(right) {
				if !visible {
					grids[i][j] = true
					result++
				}
				right = lines[i][j]
			}
		}
	}

	for i := 1; i < cols-1; i++ {
		top := lines[0][i]
		bottom := lines[rows-1][i]

		// from top to bottom
		for j := 1; j < rows-1; j++ {
			visible := grids[j][i]
			if int(lines[j][i]) > int(top) {
				if !visible {
					grids[j][i] = true
					result++
				}
				top = lines[j][i]
			}
		}

		// from bottom to top
		for j := rows - 2; j > 0; j-- {
			visible := grids[j][i]
			if int(lines[j][i]) > int(bottom) {
				if !visible {
					grids[j][i] = true
					result++
				}
				bottom = lines[j][i]
			}
		}
	}
	return result
}

func part2(input string) int {
	// parse input
	lines := strings.Split(input, "\n")
	rows, cols := len(lines), len(lines[0])

	var result int

	for i := 1; i < rows-1; i++ {
		for j := 1; j < cols-1; j++ {
			var left, right, top, bottom int
			tmp := lines[i][j]
			// the left side
			for k := j - 1; k >= 0; k-- {
				left++
				if lines[i][k] >= tmp {
					break
				}
			}
			// to the right
			for k := j + 1; k <= cols-1; k++ {
				right++
				if lines[i][k] >= tmp {
					break
				}
			}

			//to the bottom
			for k := i + 1; k <= rows-1; k++ {
				bottom++
				if lines[k][j] >= tmp {
					break
				}

			}
			// to the top
			for k := i - 1; k >= 0; k-- {
				top++
				if lines[k][j] >= tmp {
					break
				}
			}

			if left*right*top*bottom > result {
				result = left * right * top * bottom
				fmt.Println(result)
			}
		}
	}
	return result
}
