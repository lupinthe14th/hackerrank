package main

import "fmt"

func print30HelloWorld(s string) string {
	return fmt.Sprintf("Hello, World.\n%s", s)
}

func main() {
	var s string
	fmt.Scan(&s)
	fmt.Print(print30HelloWorld(s))
}
