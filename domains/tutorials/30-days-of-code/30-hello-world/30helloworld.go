package main

import (
	"bufio"
	"fmt"
	"os"
)

func print30HelloWorld(s string) string {
	return fmt.Sprintf("Hello, World.\n%s", s)
}

func main() {
	s := bufio.NewScanner(os.Stdin)
	s.Scan()
	fmt.Print(print30HelloWorld(s.Text()))
}
