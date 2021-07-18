package vowelsubstring

import (
	"fmt"
	"testing"
)

func TestFindSubstring(t *testing.T) {
	t.Parallel()
	type in struct {
		s string
		k int32
	}
	tests := []struct {
		in   in
		want string
	}{
		{in: in{s: "azerdii", k: 5}, want: "erdii"},
		{in: in{s: "qwdftr", k: 2}, want: "Not found!"},
	}
	for i, tt := range tests {
		i, tt := i, tt
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			t.Parallel()
			got := findSubstring(tt.in.s, tt.in.k)
			if got != tt.want {
				t.Fatalf("in: %v got: %v want: %v", tt.in, got, tt.want)
			}
		})
	}
}

func TestCount(t *testing.T) {
	t.Parallel()
	tests := []struct {
		in   string
		want int32
	}{
		{in: "a", want: 1},
		{in: "b", want: 0},
		{in: "aeiou", want: 5},
	}
	for i, tt := range tests {
		i, tt := i, tt
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			t.Parallel()
			got := count(tt.in)
			if got != tt.want {
				t.Fatalf("in: %v got: %v want: %v", tt.in, got, tt.want)
			}
		})
	}
}
