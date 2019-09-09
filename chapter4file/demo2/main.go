package main

import (
	"log"
	"os"
)

func main()  {
	// 裁剪文件到100个字节
	err := os.Truncate("test.txt", 10)
	if err != nil {
		log.Panicln(err)
	}

}