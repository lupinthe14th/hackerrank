package main

import (
	"reflect"
	"testing"
)

var cases = []struct {
	id    int
	input string
	want  [2]string
}{
	{id: 1, input: "Hacker", want: [2]string{"Hce", "akr"}},
	{id: 2, input: "Rank", want: [2]string{"Rn", "ak"}},
}

func TestPrintReview(t *testing.T) {
	for _, tt := range cases {
		got := printReview(tt.input)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%v, want: %v", got, tt.want)
		}
	}
}
