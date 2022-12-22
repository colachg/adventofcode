package day06

import (
	_ "embed"
	"testing"
)

//go:embed input.txt
var input string

func Test_part1(t *testing.T) {
	tests := []struct {
		name string
		args string
		want int
	}{
		// TODO: Add test cases.
		{
			"Test part1 with example: mjqjpqmgbljsphdztnvjfqwrcgsmlb",
			"mjqjpqmgbljsphdztnvjfqwrcgsmlb",
			7,
		},
		{
			"Test part1 with example: bvwbjplbgvbhsrlpgdmjqwftvncz",
			"bvwbjplbgvbhsrlpgdmjqwftvncz",
			5,
		},
		{
			"Test part1 with example: nppdvjthqldpwncqszvftbrmjlhg",
			"nppdvjthqldpwncqszvftbrmjlhg",
			6,
		},
		{
			"Test part1 with example: nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg",
			"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg",
			10,
		},
		{
			"Test part1 with example: zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw",
			"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw",
			11,
		},
		{
			"Test part1 with input",
			input,
			1760,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part1(tt.args); got != tt.want {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_part2(t *testing.T) {
	tests := []struct {
		name string
		args string
		want int
	}{
		// TODO: Add test cases.
		{
			"Test part2 with example: mjqjpqmgbljsphdztnvjfqwrcgsmlb",
			"mjqjpqmgbljsphdztnvjfqwrcgsmlb",
			19,
		},
		{
			"Test part2 with example: bvwbjplbgvbhsrlpgdmjqwftvncz",
			"bvwbjplbgvbhsrlpgdmjqwftvncz",
			23,
		},
		{
			"Test part2 with example: nppdvjthqldpwncqszvftbrmjlhg",
			"nppdvjthqldpwncqszvftbrmjlhg",
			23,
		},
		{
			"Test part2 with example: nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg",
			"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg",
			29,
		},
		{
			"Test part2 with example: zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw",
			"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw",
			26,
		},
		{
			"Test part2 with input",
			input,
			2974,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part2(tt.args); got != tt.want {
				t.Errorf("part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
