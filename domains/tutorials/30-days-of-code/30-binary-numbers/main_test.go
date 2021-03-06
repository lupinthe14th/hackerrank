package main

import (
	"fmt"
	"testing"
)

var cases = []struct {
	id    int
	input int32
	want  int
}{
	{id: 1, input: 5, want: 1},
	{id: 2, input: 13, want: 2},
	{id: 3, input: 100000, want: 2},
	{id: 4, input: 99999, want: 5},
	{id: 5, input: 439, want: 3},
	{id: 6, input: 951, want: 3},
	{id: 7, input: 1911, want: 3},
	{id: 8, input: 65535, want: 16},
	{id: 9, input: 524283, want: 16},
	{id: 10, input: 524275, want: 15},
}

func TestBinaryNumber(t *testing.T) {
	for _, tt := range cases {
		t.Run(fmt.Sprint(tt.id), func(t *testing.T) {
			got := BinaryNumber(tt.input)
			if got != tt.want {
				t.Errorf("%d, want: %d", got, tt.want)
			}
		})
	}
}
