package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func hourglassSum(matrix [][]int32) int32 {
	var sum int32
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if !(i == 1 && (j == 0 || j == 2)) {
				sum += matrix[i][j]
			}
		}
	}
	return sum
}

func hourglassMax(matrix [][]int32) int32 {
	var max int32
	for i := 0; i <= 3; i++ {
		for j := 0; j <= 3; j++ {
			m := make([][]int32, 0, 3)
			m = append(m, matrix[i][j:j+3])
			m = append(m, matrix[i+1][j:j+3])
			m = append(m, matrix[i+2][j:j+3])
			sum := hourglassSum(m)
			if max < sum {
				max = sum
			}
		}
	}
	return max
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	var arr [][]int32
	for i := 0; i < 6; i++ {
		arrRowTemp := strings.Split(readLine(reader), " ")

		var arrRow []int32
		for _, arrRowItem := range arrRowTemp {
			arrItemTemp, err := strconv.ParseInt(arrRowItem, 10, 64)
			checkError(err)
			arrItem := int32(arrItemTemp)
			arrRow = append(arrRow, arrItem)
		}

		if len(arrRow) != int(6) {
			panic("Bad input")
		}

		arr = append(arr, arrRow)
	}

	fmt.Print(hourglassMax(arr))

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
