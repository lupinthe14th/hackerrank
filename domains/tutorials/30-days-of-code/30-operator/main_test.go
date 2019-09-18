package main

import (
	"fmt"
	"testing"
)

type input struct {
	mealCost   float64
	tipPercent int32
	taxPercent int32
}

var cases = []struct {
	input input
	want  int32
}{
	{input: input{mealCost: 12.00, tipPercent: 20, taxPercent: 8}, want: 15},
	{input: input{mealCost: 15.50, tipPercent: 15, taxPercent: 10}, want: 19},
}

func TestSolve(t *testing.T) {
	for i, tt := range cases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			got := solve(tt.input.mealCost, tt.input.tipPercent, tt.input.taxPercent)
			if got != tt.want {
				t.Errorf("%d, want: %d", got, tt.want)
			}
		})
	}
}
