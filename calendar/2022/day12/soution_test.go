package day12

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
		{
			"Test part1 with example",
			example,
			31,
		},
		{
			"Test part1 with input1",
			input,
			497,
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
			"Test part1 with example",
			example,
			29,
		},
		{
			"Test part1 with input1",
			input,
			492,
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
