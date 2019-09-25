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
	{id: 1, input: 3, want: 6},
	{id: 2, input: 1, want: 1},
	{id: 3, input: 12, want: 479001600},
}

func TestFactorinal(t *testing.T) {
	for _, tt := range cases {
		t.Run(fmt.Sprint(tt.id), func(t *testing.T) {
			got := factorial(tt.input)
			if got != tt.want {
				t.Errorf("%d, want: %d", got, tt.want)
			}
		})
	}
}
