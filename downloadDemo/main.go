// 下载文件demo
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

var dirPath = ""
var destinationPath = ""

func main() {
	if len(os.Args) < 3 {
		fmt.Println("缺少参数----------")
		fmt.Scanf("输入任何字符退出---")
		return
	}

	dirPath = os.Args[1]
	destinationPath = os.Args[2]

	start := time.Now().UnixNano() / 1e6

	ch := make(chan string)
	fileList, _ := ioutil.ReadDir(dirPath)
	for _, file := range fileList {
		go copyFile(file.Name(), ch)
	}

	for i := 0; i < len(fileList); i++ {
		select {
		case str := <-ch:
			fmt.Println(i, ":", str)
		}
	}

	end := time.Now().UnixNano() / 1e6
	fmt.Println("下载结束---------", (end - start))
	fmt.Scanf("输入任何字符退出---")
}

// 复制文件到指定目录
func copyFile(name string, ch chan string) {
	input, err := ioutil.ReadFile(dirPath + "\\" + name)
	if err != nil {
		ch <- "下载失败--" + name
	}

	err = ioutil.WriteFile(destinationPath+"\\"+name, input, 0644)
	if err != nil {
		ch <- "下载失败--" + name
	}
	ch <- "下载成功--" + name
}
