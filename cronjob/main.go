package main

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/robfig/cron"
)

func main() {
	url1 := "http://192.168.1.93:8080/meinian-compliance/testLog1/index"
	url2 := "http://192.168.1.93:8080/meinian-compliance/testLog2/index"
	spec1 := "*/1 * * * * ?"
	spec2 := "*/1 * * * * ?"
	c := cron.New()
	c.AddFunc(spec1, func() {
		sendGet(url1)
	})
	c.AddFunc(spec2, func() {
		sendGet(url2)
	})
	c.Start()

	select {}
}

func sendGet(url string) {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	log.Println(string(body))
}
