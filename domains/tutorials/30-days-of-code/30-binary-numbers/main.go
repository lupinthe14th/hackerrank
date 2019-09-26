package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func BinaryNumber(n int32) int32 {
	s := fmt.Sprintf("%b", n)

	out := s[:len(s)-1]

	count := make(map[rune]int32)

	for _, v := range out {
		if v == '1' {
			count[v]++
		}
	}
	return count['1']
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
