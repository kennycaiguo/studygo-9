// studygo project main.go
package main

import (
	"log"
	"os"
	"strconv"
	"time"
)

//爆炸了
func booooom() {
	writeLog("---------booooom----------")
}

//倒计时
func countdown(litSignal chan int) {
	c := time.Tick(1 * time.Second)
	for countNum := 10; countNum > 0; countNum-- {
		writeLog("---------" + strconv.Itoa(countNum) + "---------")
		<-c
	}
	litSignal <- -1
}

//取消
func isCancel(cancelSignal chan int) {
	var buffer [256]byte
	_, err := os.Stdin.Read(buffer[:])
	if err != nil {
		panic(err)
	}
	writeLog("---------cancel---------")
	cancelSignal <- -1
}

func writeLog(msg string) {
	logFile, err := os.OpenFile("syslog.txt-"+time.Now().Format("2006-01-02"), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	defer logFile.Close()
	if err != nil {
		panic(err)
	}
	infoLog := log.New(logFile, "[INFO]", log.LstdFlags)
	infoLog.Print(msg)
}

func main() {
	writeLog("---------placing bombs---------")
	cancelSignal := make(chan int)
	litSignal := make(chan int)
	go isCancel(cancelSignal)
	go countdown(litSignal)

	select {
	case <-litSignal:
		booooom()
	case <-cancelSignal:
		return
	}
}
