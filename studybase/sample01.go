package main

import (
	"fmt"
	"os"
)

func main()  {
	filename := "/Users/lixiaoyi/GitRepo/cwowhappy/study-golang/README.md"
	if fileInfo, err := os.Stat(filename); err == nil {
		fmt.Println(fileInfo)
	}
}


