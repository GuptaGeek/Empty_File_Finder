package main

import (
	"os"
	"sort"
)

// ---------------------------------------------------------
// EXERCISE: Sort and write items to a file
//
//  1. Get arguments from command-line
//
//  2. Sort them
//
//  3. Write the sorted slice to a file
//
//
// EXPECTED OUTPUT
//
//   go run main.go
//     Send me some items and I will sort them
//
//   go run main.go orange banana apple
//
//   cat sorted.txt
//     apple
//     banana
//     orange



func main() {
	inputslice := os.Args[1:]
	// filename := inputslice[len(os.Args)-1]
	// inputslice = inputslice[1 : len(os.Args)-1]

	sort.Strings(inputslice)
	var total int
	for _, value := range inputslice {
		total += len(value)
	}

	names := make([]byte, 0, total)
	for _, value := range inputslice {
		names = append(names, value...)
	}
	os.WriteFile("nonemptyfile2", names, 0333)
	return
}
