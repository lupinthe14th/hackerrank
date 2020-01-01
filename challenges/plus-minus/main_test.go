package main

import (
	"fmt"
	"reflect"
	"testing"
)

var cases = []struct {
	id    int
	input []int32
	want  Pm
}{
	{id: 1, input: []int32{-4, 3, -9, 0, 4, 1}, want: Pm{p: 0.500000, n: 0.3333333333333333, z: 0.16666666666666666}},
	{id: 2, input: []int32{1, 2, 3, -1, -2, -3, 0, 0}, want: Pm{p: 0.375000, n: 0.375000, z: 0.250000}},
}

func Example_plusMinus() {
	plusMinus(cases[0].input)

	// Output:
	// 0.500000
	// 0.333333
	// 0.166667
}

func TestSubPlusMinus(t *testing.T) {
	for _, tt := range cases {
		t.Run(fmt.Sprint(tt.id), func(t *testing.T) {
			got := subPlusMinus(tt.input)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("%v ,want: %v", got, tt.want)
			}
		})
	}
}
