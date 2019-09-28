package main

import (
	"fmt"
	"testing"
)

func TestHourglassMax(t *testing.T) {
	var cases = []struct {
		id    int
		input [][]int32
		want  int32
	}{{id: 1,
		input: [][]int32{
			{1, 1, 1, 0, 0, 0},
			{0, 1, 0, 0, 0, 0},
			{1, 1, 1, 0, 0, 0},
			{0, 0, 2, 4, 4, 0},
			{0, 0, 0, 2, 0, 0},
			{0, 0, 1, 2, 4, 0},
		},
		want: 19},
		{id: 2, input: [][]int32{
			{1, 1, 1, 0, 0, 0},
			{0, 1, 0, 0, 0, 0},
			{1, 1, 1, 0, 0, 0},
			{0, 9, 2, -4, -4, 0},
			{0, 0, 0, -2, 0, 0},
			{0, 0, -1, -2, -4, 0},
		}, want: 13},
		{id: 3, input: [][]int32{
			{-1, -1, 0, -9, -2, -2},
			{-2, -1, -6, -8, -2, -5},
			{-1, -1, -1, -2, -3, -4},
			{-1, -9, -2, -4, -4, -5},
			{-7, -3, -3, -2, -9, -9},
			{-1, -3, -1, -2, -4, -5},
		}, want: -6},
	}
	for _, tt := range cases {
		t.Run(fmt.Sprint(tt.id), func(t *testing.T) {
			got := hourglassMax(tt.input)
			if got != tt.want {
				t.Errorf("%d, want: %d", got, tt.want)
			}
		})
	}
}

func TestHourglassSum(t *testing.T) {
	var cases = []struct {
		id    int
		input [][]int32
		want  int32
	}{
		{id: 1, input: [][]int32{{1, 1, 1},
			{0, 1, 0},
			{1, 1, 1}}, want: 7},
		{id: 2, input: [][]int32{{2, 4, 4},
			{0, 2, 0},
			{1, 2, 4}}, want: 19},
		{id: 3, input: [][]int32{{0, 1, 0},
			{1, 1, 1},
			{0, 9, 2}}, want: 13},
	}
	for _, tt := range cases {
		t.Run(fmt.Sprint(tt.id), func(t *testing.T) {
			got := hourglassSum(tt.input)
			if got != tt.want {
				t.Errorf("%d, want: %d", got, tt.want)
			}
		})
	}
}
