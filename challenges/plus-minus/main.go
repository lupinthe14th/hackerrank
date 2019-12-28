package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Pm is plusMinus answer structure.
type Pm struct {
	p float64 // positive
	n float64 // negative
	z float64 // zero
}

func subPlusMinus(arr []int32) Pm {
	var p, n, z float64
	l := float64(len(arr))

	for _, v := range arr {
		switch {
		case v == 0:
			z++
		case v > 0:
			p++
		case v < 0:
			n++
		}
	}
	pm := Pm{p: p / l, n: n / l, z: z / l}

	return pm
}

// Complete the plusMinus function below.
func plusMinus(arr []int32) {
	pm := subPlusMinus(arr)
	fmt.Printf("%.6f\n", pm.p)
	fmt.Printf("%.6f\n", pm.n)
	fmt.Printf("%.6f\n", pm.z)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	n := int32(nTemp)

	arrTemp := strings.Split(readLine(reader), " ")

	var arr []int32

	for i := 0; i < int(n); i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int32(arrItemTemp)
		arr = append(arr, arrItem)
	}

	plusMinus(arr)
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
