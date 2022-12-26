package day08

import (
	_ "embed"
	"testing"
)

var (
	//go:embed example.txt
	example string
	//go:embed input.txt
	input string
)

func Test_part1(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		// TODO: Add test cases.
		{
			"Test part1 with example",
			example,
			21,
		},
		{
			"Test part1 with input",
			input,
			1705,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part1(tt.input); got != tt.want {
				t.Errorf("part2() = %v, want %v", got, tt.want)
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
		// TODO: Add test cases.
		{
			"Test part2 with example",
			example,
			8,
		},
		{
			"Test part2 with input",
			input,
			371200,
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
