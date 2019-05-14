package main

import (
	"fmt"
	"net/http"
)

var maxWorkers = 20

// JobQueue 作业队列
var JobQueue = make(chan Job, maxWorkers)

func init() {
	InitPool()
}

func main() {
	http.HandleFunc("/test/pool/", indexHandler)
	http.ListenAndServe(":9000", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	service := MyService{}
	//service.WriteInfo()
	work := Job{MyService: service}
	JobQueue <- work
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	result := "{\"msg\":\"SUCCESS\",\"code\":0}"
	fmt.Fprintln(w, result)
}
