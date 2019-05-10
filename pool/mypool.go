package main

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
				job.MyService.WriteInfo()
			}
		}
	}()
}
