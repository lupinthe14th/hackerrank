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
			got := print30HelloWorld(tt.input)
			if got != tt.want {
				t.Errorf("%s, want: %s", got, tt.want)
			}
		})
	}
}
