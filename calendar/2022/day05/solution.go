package day05

import (
	_ "embed"
	"strconv"
	"strings"
)

type Operation struct {
	Qty, From, To int
}

func parse(input string) (map[int][]byte, []Operation) {
	var row int

	// separate crates and operations
	lines := strings.Split(input, "\n")
	for n, line := range lines {
		if line == "" {
			row = n - 1 // line number of stacks id "1  2  3"
			break
		}
	}
	// parse crates
	seq := strings.Split(lines[row], "  ")
	stacks := make(map[int][]byte, len(seq))

	for i := row - 1; i >= 0; i-- {
		// [Z] [M] [P]
		for j, char := range []byte(lines[i]) {
			if char >= 'A' && char <= 'Z' {
				index := j/4 + 1 // index divided by 4.
				stacks[index] = append(stacks[index], char)
			}
		}
	}

	// parse operationss
	var operations = make([]Operation, 0)
	for i, j := row+2, 0; i < len(lines); i, j = i+1, j+1 {
		if lines[i] != "" {
			line := strings.Split(lines[i], " ") //be careful of splitting
			from, _ := strconv.Atoi(line[3])
			to, _ := strconv.Atoi(line[5])
			qty, _ := strconv.Atoi(line[1])
			operations = append(operations, Operation{qty, from, to})
		}
	}
	return stacks, operations
}

func getTopCrate(stacks map[int][]byte) string {
	var result string
	for i := 1; i <= len(stacks); i++ {
		last := len(stacks[i]) - 1
		result += string(stacks[i][last])
	}
	return result
}

func part1(input string) string {
	stacks, operations := parse(input)

	for _, opt := range operations {
		for i := 0; i < opt.Qty; i++ {
			//pop
			last := len(stacks[opt.From]) - 1
			if last < 0 {
				break
			}
			elem := stacks[opt.From][last]
			stacks[opt.From] = stacks[opt.From][:last]

			//push
			stacks[opt.To] = append(stacks[opt.To], elem)
		}
	}
	return getTopCrate(stacks)
}

func part2(input string) string {
	stacks, operations := parse(input)

	for _, opt := range operations {
		first := len(stacks[opt.From]) - opt.Qty
		last := len(stacks[opt.From])
		stacks[opt.To] = append(stacks[opt.To], stacks[opt.From][first:last]...)
		stacks[opt.From] = stacks[opt.From][:first]
	}
	return getTopCrate(stacks)
}
