package day04

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

func Test_puzzle(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
		want2 int
	}{
		{
			"Test part1 with example",
			example,
			2, 4,
		},
		{
			"Test part1 with input",
			input,
			471, 888,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, got2 := puzzle(tt.input); got != tt.want && got != tt.want2 {
				t.Errorf("got1 = %v, want1 %v, got2 = %v, want2 %v", got, tt.want, got2, tt.want2)
			}
		})
	}
}
