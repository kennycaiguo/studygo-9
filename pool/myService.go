package main

import (
	"log"
	"os"
	"strconv"
	"time"
)

// MyService 业务接口
type MyService struct {
}

// WriteInfo 写日志
func (s *MyService) WriteInfo() {
	time.Sleep(1 * time.Second)
	t := time.Now()
	logFile, err := os.OpenFile("syslog.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	defer logFile.Close()
	if err != nil {
		panic(err)
	}
	infoLog := log.New(logFile, "[INFO]", log.LstdFlags)
	infoLog.Print("time=" + strconv.FormatInt(t.UTC().UnixNano(), 10))
}
