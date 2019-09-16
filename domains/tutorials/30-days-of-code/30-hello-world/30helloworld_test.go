package main

import (
	"fmt"
	"testing"
)

var cases = []struct {
	input string
	want  string
}{
	{input: "Welcome to 30 Days of Code!", want: "Hello, World.\nWelcome to 30 Days of Code!"},
}

func TestPrint30HelloWorld(t *testing.T) {
	for _, tt := range cases {
		t.Run(fmt.Sprint(tt.input), func(t *testing.T) {
			actual := print30HelloWorld(tt.input)
			if actual != tt.want {
				t.Errorf("%s, want: %s", actual, tt.want)
			}
		})
	}
}
