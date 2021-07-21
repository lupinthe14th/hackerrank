package main

import (
	"os"
	"path/filepath"
)

func Example_main() {
	fd, _ := os.Open(filepath.Join("testdata", "input.txt"))
	orgStdin := os.Stdin
	os.Stdin = fd
	defer func() {
		os.Stdin = orgStdin
		fd.Close()
	}()

	main()

	// Output:
	// 1
	// 3
	// 5
	// 7
	// 9
	// 2
	// 4
	// 6
	// 8
	// 10
}

// func Example_main() {
// 	fd, _ := os.Open(filepath.Join("testdata", "input1.txt"))
// 	orgStdin := os.Stdin
// 	os.Stdin = fd
// 	defer func() {
// 		os.Stdin = orgStdin
// 		fd.Close()
// 	}()
//
// 	main()
//
// 	// Output:
// 	// 447
// 	// -543
// 	// 238
// 	// 122
// 	// 200
// 	// 0
// }
