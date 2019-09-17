package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var _ = strconv.Itoa // Ignore this comment. You can still use the package "strconv".
	var i uint64 = 4
	var d float64 = 4.0
	var s string = "HackerRank "

	scanner := bufio.NewScanner(os.Stdin)
	// Declare second integer, double, and String variables.
	var si uint64
	var sd float64
	var ss string

	// Read and save an integer, double, and String to your variables
	scanner.Scan()
	si, err := strconv.ParseUint(scanner.Text(), 10, 64)
	if err != nil {
		panic(err)
	}
	scanner.Scan()
	sd, err = strconv.ParseFloat(scanner.Text(), 64)
	if err != nil {
		panic(err)
	}
	scanner.Scan()
	ss = scanner.Text()

	// Print the sum of both integer variables on a new line.
	fmt.Println(i + si)
	// Print the sum of the double variables on a new line.
	fmt.Printf("%0.1f\n", d+sd)
	// Concatenate and print the String variables on a new line
	// The 's' variable above should be printed first.
	fmt.Println(s + ss)
}
