package main

import (
	"log"
	"os"
)

var (
	newFile *os.File
	err error
)
func main()  {
	// 创建一个文件
	newFile, err = os.Create("test.txt")
	if err != nil {
		log.Fatal(err)
	}

	log.Println(newFile)
	newFile.Close()
}