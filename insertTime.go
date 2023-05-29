package main

import (
	"os"
	"time"
)

func insertTime() {
	fileName := "../../dataset/data.txt"
	content := ""
	f, err := os.OpenFile(fileName, os.O_WRONLY, 0644)
	if err != nil {
		// 打开文件失败处理

	} else {
		content = "\r\n" + time.Now().Local().String()

		// 查找文件末尾的偏移量
		n, _ := f.Seek(0, 2)

		// 从末尾的偏移量开始写入内容
		_, err = f.WriteAt([]byte(content), n)
	}

	defer f.Close()
	/*
		file, err := os.Create("data.txt")
		if err != nil {
			fmt.Println(err)
		}

		defer file.Close()

		file.WriteString(time.Now().Local().String())
	*/
}
