package day13

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
			"Test with example",
			example,
			13,
		},

		{
			"Test with input",
			input,
			5506,
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
		// TODO: Add test cases.
		{
			"Test with example",
			example,
			140,
		},
		{
			"Test with input",
			input,
			21756,
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
