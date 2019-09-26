package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

func BinaryNumber(n int32) int {
	s := fmt.Sprintf("%b", n)

	var stack []int

	c := 0
	for i, v := range s {
		if v == '1' {
			c++
			if i == len(s)-1 {
				stack = append(stack, c)
			}
		} else {
			stack = append(stack, c)
			c = 0
		}
	}
	sort.Ints(stack)
	return stack[len(stack)-1]
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	n := int32(nTemp)
	fmt.Print(BinaryNumber(n))
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
