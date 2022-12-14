package day11

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
			name:  "Test part1 with example after round 20",
			input: example,
			want:  10605,
		},

		{
			name:  "Test part1 with input after round 20",
			input: input,
			want:  57348,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part1(tt.input, 20); got != tt.want {
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
			name:  "Test part1 with example after round 10000",
			input: example,
			want:  2713310158,
		},
		{
			name:  "Test part1 with input after round 20",
			input: input,
			want:  14106266886,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part2(tt.input, 10000); got != tt.want {
				t.Errorf("part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
