package helper

import (
	"bufio"
	"fmt"
	"os"
)

type Enter struct {
	Id            int
	Enter_key     string
	Enter_type    int
	Enter_name    string
	Is_next_month int
}

func WriteToFile(fileName string, clients []Enter) (bool, error) {

	outPath := "/Users/gongyao/workspace/goproject/src/mygo2/"

	// 检查路径
	checkPath(outPath)
	// 构造输出文件
	filePath := outPath + fileName

	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return false, err
	}
	//及时关闭file句柄
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic("error")
		}
	}(file)
	//写入文件时，使用带缓存的 *Writer
	write := bufio.NewWriter(file)

	_, err = write.WriteString(fmt.Sprint("来源", "\t", "客户名称", "\t", "二月份有没有下单", "\n"))
	if err != nil {
		return false, err
	}

	for _, client := range clients {
		var enTypeName string
		if client.Enter_type == 1 {
			enTypeName = "箱信中台客户"
		} else {
			enTypeName = "Erp"
		}
		_, err := write.WriteString(fmt.Sprint(enTypeName, "\t", client.Enter_name, "\t", client.Is_next_month, "\n"))
		if err != nil {
			return false, err
		}
	}

	err = write.Flush()
	if err != nil {
		return false, err
	}

	return true, nil
}

func checkPath(outPath string) bool {
	_, err := os.Stat(outPath)
	fmt.Println(outPath)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		panic("文件夹不存在")
	}
	fmt.Println(os.IsTimeout(err))

	panic(err.Error())

	return false
}
