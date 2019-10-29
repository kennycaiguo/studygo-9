package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type resultMsg struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
}

func uploadHandler(response http.ResponseWriter, request *http.Request) {
	// 获取文件参数
	file, fileHeader, err := request.FormFile("file")
	if err != nil {
		returnMsg(response, "-1", err.Error())
		return
	}
	defer file.Close()

	fmt.Println("-----接收到的文件：", fileHeader.Filename, ",文件大小：", fileHeader.Size, "------")

	// 创建文件
	newFile, err := os.Create("E://" + fileHeader.Filename)
	if err != nil {
		returnMsg(response, "-1", err.Error())
		return
	}
	defer newFile.Close()

	// 保存文件到本地
	_, err = io.Copy(newFile, file)
	if err != nil {
		returnMsg(response, "-1", err.Error())
		return
	}

	returnMsg(response, "0", "success")
}

func main() {
	http.HandleFunc("/upload", uploadHandler)
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		fmt.Println("启动失败", err)
	}
}

func returnMsg(response http.ResponseWriter, code string, msg string) {
	msgJSON := &resultMsg{code, msg}
	b, err := json.Marshal(msgJSON)
	if err != nil {
		fmt.Fprintln(response, "error parse msg")
		return
	}
	fmt.Fprintln(response, string(b))
}
