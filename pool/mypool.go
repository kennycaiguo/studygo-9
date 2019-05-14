package main

// Job 表示要执行的作业
type Job struct {
	MyService MyService
}

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

// InitPool 给池中初始化一定量的工人，以及开启任务队列的监听
func InitPool() {
	// 创建工作池
	pool := make(chan chan Job, maxWorkers)
	// 创建一定数量的工人（可以看做：创建了N个工人，每个工人能并发处理N件工作）
	for i := 0; i < maxWorkers; i++ {
		worker := NewWorker(pool)
		worker.Start()
	}
	//监听JobQueue上是否有新任务
	go func() {
		for {
			select {
			case job := <-JobQueue:
				go func(job Job) {
					// 获取可用的工人channel，若没有，则阻塞
					jobChannel := <-pool
					jobChannel <- job
				}(job)
			}
		}
	}()
}
