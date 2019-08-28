package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	f, err := os.Open("CTolua.log")
	if err != nil {
		panic(err)
	}
	fmt.Println("开始读取------------")
	buff := bufio.NewReader(f)
	for {
		line, err := buff.ReadString('\n')
		if err != nil || io.EOF == err {
			break
		}
		res := sliceStr(line)
		if res != "" {
			writelog("INFO", res+"\r\n")
		}
	}
	fmt.Println("读取结束------------")
}

func sliceStr(str string) string {
	if str != "" {
		indexMethod := strings.Index(str, "&reviewTime=2019-06-11")
		//indexStatus := strings.Index(str, "MANUAL_CHECKED")
		indexAdmIDIss := strings.Index(str, "&admIdIss=2&applyParam=")

		if indexMethod != -1 && indexAdmIDIss != -1 {
			return str
		}
	}
	return ""
}

func writelog(logGrade string, msg string) {
	logFile, err := os.OpenFile("log.txt-"+time.Now().Format("2006-01-02"), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	defer logFile.Close()
	if err != nil {
		fmt.Println("open log file error")
	}
	infoLog := log.New(logFile, "["+logGrade+"]", log.LstdFlags)
	infoLog.Print(msg)
}
