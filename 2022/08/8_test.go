package main

import (
	"testing"
)

func Test_part1(t *testing.T) {
	type args struct {
		input [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example",
			args: args{
				input: [][]int{
					{3, 0, 3, 7, 3},
					{2, 5, 5, 1, 2},
					{6, 5, 3, 3, 2},
					{3, 3, 5, 4, 9},
					{3, 5, 3, 9, 0},
				},
			},
			want: 21,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part1(tt.args.input); got != tt.want {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_part2(t *testing.T) {
	type args struct {
		input [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example",
			args: args{
				input: [][]int{
					{3, 0, 3, 7, 3},
					{2, 5, 5, 1, 2},
					{6, 5, 3, 3, 2},
					{3, 3, 5, 4, 9},
					{3, 5, 3, 9, 0},
				},
			},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part2(tt.args.input); got != tt.want {
				t.Errorf("part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
