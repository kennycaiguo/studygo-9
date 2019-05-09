package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

var maxWorkers = 20

// MyService 业务接口
type MyService struct {
}

// Job 表示要执行的作业
type Job struct {
	MyService MyService
}

// JobQueue 作业队列
var JobQueue chan Job

// Worker 执行作业的工人
type Worker struct {
	WorkerPool chan chan Job
	JobChannel chan Job
}

// NewWorker 新建一个工人
func NewWorker(workerPool chan chan Job) Worker {
	return Worker{
		WorkerPool: workerPool,
		JobChannel: make(chan Job),
	}
}

// Start 监听
func (w Worker) Start() {
	go func() {
		for {
			w.WorkerPool <- w.JobChannel
			select {
			case job := <-w.JobChannel:
				//有工作任务时，开始执行业务接口的方法
				job.MyService.writeInfo()
			}
		}
	}()
}

func init() {
	// 创建工人池
	//pool := make(chan chan Job, maxWorkers)
	// 创建一定数量的工人
	for i := 0; i < maxWorkers; i++ {

	}

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
