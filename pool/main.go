package main

import (
	"fmt"
	"net/http"
)

var maxWorkers = 20

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
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	result := "{\"msg\":\"SUCCESS\",\"code\":0}"
	fmt.Fprintln(w, result)
}
