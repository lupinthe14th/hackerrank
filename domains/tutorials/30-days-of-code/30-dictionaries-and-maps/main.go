package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var entryNumber, phoneNumber int
	var name, queryName string
	fmt.Scan(&entryNumber)

	phonebook := make(map[string]int)
	for i := 0; i < entryNumber; i++ {
		fmt.Scan(&name, &phoneNumber)
		phonebook[name] = phoneNumber
	}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		queryName = scanner.Text()
		number, ok := phonebook[queryName]
		if ok {
			fmt.Printf("%s=%d\n", queryName, number)
		} else {
			fmt.Println("Not found")
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}
