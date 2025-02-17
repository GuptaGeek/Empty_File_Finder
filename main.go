package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {

	path := "C:\\Users\\user\\OneDrive\\Desktop\\Directory"
	files, err := ioutil.ReadDir(path)
	var total int
	for _, file := range files {
		if file.Size() == 0 {
			total += len(file.Name()) + 1
		}
	}

	if err != nil {
		fmt.Println("ERROR")
		return
	}

	var names = make([]byte, 0, total)
	for _, file := range files {
		if file.Size() == 0 {
			name := file.Name()
			names = append(names, name...)
			names = append(names, '\n')
		}
	}

	//The type of files slice each value is os.FileInfo
	//FileInfo is type of interface in go lang
	//The struct associated with it is the os.fileStat
	os.WriteFile("nonemptyfile1", names, 755)
	fmt.Printf("%s\n", names)

}
