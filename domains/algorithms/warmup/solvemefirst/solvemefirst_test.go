package main

import (
	"fmt"
	"testing"
)

var cases = []struct {
	a, b uint32
	want uint32
}{
	{a: 2, b: 3, want: 5},
}

func TestSolveMeFirst(t *testing.T) {
	for _, tt := range cases {
		t.Run(fmt.Sprintf("a: %d, b: %d", tt.a, tt.b), func(t *testing.T) {
			actual := solveMeFirst(tt.a, tt.b)
			if actual != tt.want {
				t.Errorf("%d, want: %d", actual, tt.want)
			}
		})
	}
}
