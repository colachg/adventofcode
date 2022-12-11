package day01

import (
	_ "embed"
	"reflect"
	"testing"
)

var (
	//go:embed input.txt
	input []byte
	//go:embed example.txt
	example1 []byte
)

func Test_part1(t *testing.T) {
	tests := []struct {
		name  string
		input []byte
		want  int
	}{
		{
			"Test part1 with example",
			example1,
			24000,
		},
		{
			"Test part1 with my input",
			input,
			71502,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part1(tt.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addTwoNumbers() = %v, want %v", got, tt.want)
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
			example1,
			45000,
		},
		{
			"Test part2 with my input",
			input,
			208191,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			part2(tt.input)
		})
	}
}
