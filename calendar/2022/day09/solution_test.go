package main

import (
	_ "embed"
	"testing"
)

var (
	//go:embed input.txt
	input string

	//go:embed example.txt
	example string

	//go:embed example2.txt
	example2 string
)

func Test_part1(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{"Test part1 with example",
			example,
			13,
		},
		{
			"Test part1 with input",
			input,
			6498,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part1(tt.input); got != tt.want {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_part2(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			"Test part2 with example",
			example2,
			36,
		},
		{
			"Test part2 with input",
			input,
			2531,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part2(tt.input); got != tt.want {
				t.Errorf("part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
