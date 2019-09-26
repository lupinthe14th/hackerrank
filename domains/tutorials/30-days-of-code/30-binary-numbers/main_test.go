package main

import (
	"fmt"
	"testing"
)

var cases = []struct {
	id    int
	input int32
	want  int32
}{
	{id: 1, input: 5, want: 1},
	{id: 2, input: 13, want: 2},
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
