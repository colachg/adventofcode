package day03

import (
	_ "embed"
	"testing"
)

var (
	//go:embed input.txt
	input []byte
	//go:embed example.txt
	example []byte
)

func Test_part1(t *testing.T) {
	tests := []struct {
		name  string
		input []byte
		want  int
	}{
		{
			"Test part1 with example",
			example,
			157,
		},
		{
			"Test part1 with input",
			input,
			7889,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part1(tt.input); got != tt.want {
				t.Errorf("got = %v, want %v", got, tt.want)
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
			70,
		},
		{
			"Test part2 with input",
			input,
			2825,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part2(tt.input); got != tt.want {
				t.Errorf("got = %v, want %v", got, tt.want)
			}
		})
	}
}
