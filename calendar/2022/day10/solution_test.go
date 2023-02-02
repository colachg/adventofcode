package day10

import (
	_ "embed"
	"reflect"
	"testing"
)

var (

	//go:embed example.txt
	example string

	//go:embed input.txt
	input string
)

func Test_parse(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  []int
	}{
		// TODO: Add test cases.
		{
			"Test parse with example",
			"noop\naddx 3\naddx -5",
			[]int{1, 1, 1, 4, 4, -1},
		},
		{
			"Test parse with example2",
			"noop\naddx 3\naddx -5\naddx 6\n noop",
			[]int{1, 1, 1, 4, 4, -1, -1, 5, 5},
		},
		{
			"Test parse with example3",
			"noop\nnoop\nnoop\nnoop",
			[]int{1, 1, 1, 1, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parse(tt.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parse() = %v, want %v", got, tt.want)
			}
		})
	}
}

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
			13140,
		},
		{
			"Test part1 with input",
			input,
			15220,
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
			"Test part2 with example",
			example,
			0,
		},
		{
			"Test part2 with input",
			input,
			0,
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
