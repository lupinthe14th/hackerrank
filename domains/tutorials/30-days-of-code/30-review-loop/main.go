package main

import (
	"fmt"
	"strings"
)

func printReview(s string) [2]string {
	odd := make([]string, 0)
	even := make([]string, 0)
	for i, v := range strings.Split(s, "") {
		if i%2 == 0 {
			odd = append(odd, v)
		} else {
			even = append(even, v)
		}
	}
	r := [2]string{strings.Join(odd, ""), strings.Join(even, "")}
	return r
}

func main() {

	var i int
	var s string
	fmt.Scan(&i)
	r := make([][2]string, 0, i)
	for i > 0 {
		fmt.Scan(&s)
		r = append(r, printReview(s))
		i--
	}

	for _, v := range r {
		fmt.Printf("%s %s\n", v[0], v[1])
	}
}
