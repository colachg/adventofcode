package day02

import (
	_ "embed"
	"testing"
)

//go:embed example.txt
var example []byte

//go:embed input.txt
var input []byte

func Test_part1(t *testing.T) {

	tests := []struct {
		name  string
		input []byte
		want  int
	}{
		{
			"Test part1 with example",
			example,
			15,
		},
		{
			"Test part1 with input",
			input,
			13682,
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
		input []byte
		want  int
	}{
		{
			"Test part2 with example",
			example,
			12,
		},
		{
			"Test part2 with input",
			input,
			12881,
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
