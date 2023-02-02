package day07

import (
	_ "embed"
	"reflect"
	"testing"
)

var (
	//go:embed input.txt
	input string

	//go:embed example.txt
	example string
)

func Test_parse(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  map[string]int
	}{
		{
			"Test parse with example",
			example,
			map[string]int{
				"/":    48381165,
				"/a":   94853,
				"/a/e": 584,
				"/d":   24933642,
			},
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
		name string
		fs   map[string]int
		want int
	}{
		{
			"Test part1 with example",
			parse(example),
			95437,
		},
		{
			"Test part1 with input",
			parse(input),
			1077191,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part1(tt.fs); got != tt.want {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_part2(t *testing.T) {
	tests := []struct {
		name string
		fs   map[string]int
		want int
	}{
		{
			"Test part2 with example",
			parse(example),
			24933642,
		},
		{
			"Test part2 with input",
			parse(input),
			5649896,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part2(tt.fs); got != tt.want {
				t.Errorf("part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
