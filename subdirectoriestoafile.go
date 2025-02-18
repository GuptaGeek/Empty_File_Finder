package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

// ---------------------------------------------------------
// EXERCISE: Find and write the names of subdirectories to a file
//
//  Create a program that can get multiple directory paths from
//  the command-line, and prints only their subdirectories into a
//  file named: dirs.txt
//
//
//  1. Get the directory paths from command-line
//
//  2. Append the names of subdirectories inside each directory
//     to a byte slice
//
//  3. Write that byte slice to dirs.txt file
//
//
// EXPECTED OUTPUT
//
//   go run main.go
//     Please provide directory paths
//
//   go run main.go dir/ dir2/
//
//   cat dirs.txt
//
//     dir/
//             subdir1/
//             subdir2/
//
//     dir2/
//             subdir1/
//             subdir2/
//             subdir3/

func main() {
	paths := os.Args[1:]
	if len(paths) == 1 {
		fmt.Println("please enter directoreis")
	}
	var dirs []byte
	for _, path := range paths {
		files, err := ioutil.ReadDir(path)
		if err != nil {
			fmt.Println("ERROR")
			return
		}
		dirs = append(dirs, path...)
		dirs = append(dirs, '\\')
		for _, file := range files {
			if file.IsDir() == true {
				dirs = append(dirs, '\t')
				dirs = append(dirs, file.Name()...)
				dirs = append(dirs, '\\', '\n')
			}
		}
	}
	os.WriteFile("dirs.text", dirs, 0644)
}
