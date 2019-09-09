package main

import (
	"fmt"
	"os"
)

func main()  {
	// 打开文件
	file, err := os.OpenFile("test.txt", os.O_WRONLY, 0666)
	if err != nil {
		switch {
		case os.IsPermission(err):
			fmt.Println("Permission Error")
			break
		case os.IsNotExist(err):
			fmt.Println("Not Exist")
			break
		default:
			fmt.Println(err)
		}
	}
	defer file.Close()


}