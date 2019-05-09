package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

// MyService 业务接口
type MyService struct {
}

func main() {
	http.HandleFunc("/test/pool/", indexHandler)
	http.ListenAndServe(":9000", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	service := MyService{}
	go service.writeInfo()
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	result := "{\"msg\":\"SUCCESS\",\"code\":0}"
	fmt.Fprintln(w, result)
}

func (s *MyService) writeInfo() {
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
