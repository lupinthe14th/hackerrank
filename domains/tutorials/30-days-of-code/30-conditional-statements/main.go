package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func printConditionalAction(n int32) (string, error) {
	// Constraints check
	if !(n >= 1) && !(n <= 100) {
		return "", errors.New("n is constraints are not satisfied")
	}
	// n is odd
	if n%2 == 1 {
		return "Weird", nil
	}
	// n is even
	if n%2 == 0 {
		// And n is in the inclusive range of 2 to 5
		if (n >= 2) && (n <= 5) {
			return "Not Weird", nil
		}
		// And n is in the inclusive range of 6 to 20
		if (n >= 6) && (n <= 20) {
			return "Weird", nil
		}
		if n > 20 {
			return "Not Weird", nil
		}
	}
	return "", errors.New("n is constraints are not satisfied")
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	NTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	N := int32(NTemp)
	printString, err := printConditionalAction(N)
	checkError(err)
	fmt.Print(printString)
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
