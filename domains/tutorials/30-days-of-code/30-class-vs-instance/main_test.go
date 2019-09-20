package main

import (
	"fmt"
	"testing"
)

func TestPerson(t *testing.T) {
	var cases = []struct {
		caseID    int
		input     int
		want      int
		want3yago int
	}{
		{caseID: 0, input: -1, want: 0, want3yago: 3},
		{caseID: 1, input: 10, want: 10, want3yago: 13},
		{caseID: 2, input: 16, want: 16, want3yago: 19},
		{caseID: 3, input: 18, want: 18, want3yago: 21},
	}
	for _, tt := range cases {
		got := person{age: tt.input}
		t.Run(fmt.Sprint(tt.caseID), func(t *testing.T) {
			got = got.NewPerson(tt.input)
			if got.age != tt.want {
				t.Errorf("%d, want: %d", got.age, tt.want)
			}
		})
		t.Run(fmt.Sprint(tt.caseID), func(t *testing.T) {
			for i := 0; i < 3; i++ {
				got = got.yearPasses()
			}
			if got.age != tt.want3yago {
				t.Errorf("%d, want: %d", got.age, tt.want3yago)
			}
		})
	}
}
