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

	// Output: {"full_name":"Jack Oliver","position":"CEO","age":44}
}

// func Example_main() {
// 	fd, _ := os.Open(filepath.Join("testdata", "input.txt"))
// 	orgStdin := os.Stdin
// 	os.Stdin = fd
// 	defer func() {
// 		os.Stdin = orgStdin
// 		fd.Close()
// 	}()
//
// 	main()
//
// 	// Output:{"full_name":"Jack Oliver","age":44,"years_in_company":8}
// }
