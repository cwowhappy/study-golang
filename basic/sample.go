package basic

import (
	"bytes"
	"fmt"
)

func StudySlice() {
	path := []byte("AAA/BBB")
	sepIndex := bytes.IndexByte(path, '/')

	path1 := path[:sepIndex]
	path2 := path[sepIndex+1:]
	fmt.Println(string(path1))
	fmt.Println(string(path2))

	path1 = append(path1, "SSS"...)
	fmt.Println("er append")
	fmt.Println(string(path1))
	fmt.Println(string(path2))


	path = []byte("AAA/BBB")
	sepIndex = bytes.IndexByte(path, '/')

	path1 = path[:sepIndex:sepIndex]
	path2 = path[sepIndex+1:]
	fmt.Println(string(path1))
	fmt.Println(string(path2))

	path1 = append(path1, "SSS"...)
	fmt.Println("er append")
	fmt.Println(string(path1))
	fmt.Println(string(path2))
}