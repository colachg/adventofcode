package day05

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
		want  map[int][]byte
		want1 []Operation
	}{
		{
			name:  "Test parse with example",
			input: example,
			want: map[int][]byte{
				1: {'Z', 'N'},
				2: {'M', 'C', 'D'},
				3: {'P'},
			},
			want1: []Operation{
				{1, 2, 1},
				{3, 1, 3},
				{2, 2, 1},
				{1, 1, 2},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := parse(tt.input)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parse() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("parse() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_part1(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "Test part1 with example.",
			input: example,
			want:  "CMZ",
		},
		{
			name:  "Test part1 with input.",
			input: input,
			want:  "MQTPGLLDN",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part1(tt.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got: %v, want: %v", string(got), string(tt.want))
			}
		})
	}
}

func Test_part2(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "Test part2 with example.",
			input: example,
			want:  "MCD",
		},
		{
			name:  "Test part2 with input.",
			input: input,
			want:  "LVZPSTTCZ",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part2(tt.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got = %v, want %v", string(got), string(tt.want))
			}
		})
	}
}

func Test_getTopCrate(t *testing.T) {
	tests := []struct {
		name string
		args map[int][]byte
		want string
	}{
		{
			"Test getTopCrate case 1",
			map[int][]byte{
				1: {'Z', 'N'},
				2: {'M', 'C', 'D'},
				3: {'P'},
			},
			"NDP",
		},
		{
			"Test getTopCrate case 2",
			map[int][]byte{
				1: {'Z', 'A'},
				2: {'M', 'C', 'B'},
				3: {'P'},
				4: {'K'},
			},
			"ABPK",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getTopCrate(tt.args); got != tt.want {
				t.Errorf("getTopCrate() = %v, want %v", got, tt.want)
			}
		})
	}
}
