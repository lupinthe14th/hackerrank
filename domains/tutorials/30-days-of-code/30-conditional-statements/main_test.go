package main

import (
	"fmt"
	"testing"
)

var cases = []struct {
	input  int32
	want   string
	errstr string
}{
	{input: 3, want: "Weird"},
	{input: 24, want: "Not Weird"},
	{input: 24, want: "Not Weird"},
}

func TestPrintConditionalAction(t *testing.T) {
	for _, tt := range cases {
		t.Run(fmt.Sprint(tt.input), func(t *testing.T) {
			got, err := printConditionalAction(tt.input)
			if err != nil && tt.errstr == "" {
				if got != tt.want {
					t.Errorf("%s, want: %s", got, tt.want)
				}
			}
		})
	}
}
