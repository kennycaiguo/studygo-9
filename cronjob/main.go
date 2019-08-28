package main

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"

	"github.com/robfig/cron"
)

func main() {
	url1 := "http://192.168.1.93:8080/amol-back/test/upload"
	spec1 := "*/1 * * * * ?"
	c := cron.New()
	c.AddFunc(spec1, func() {
		postFile(url1)
	})
	// c.AddFunc(spec2, func() {
	// 	sendGet(url2)
	// })
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

func postFile(url string) {
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)
	fileWriter, err := bodyWriter.CreateFormFile("img-blob", "JCBG190411025747799878.pdf")
	if err != nil {
		panic(err)
	}

	file, err := os.Open("E:/JCBG190411025747799878.pdf")
	if err != nil {
		panic(err)
	}

	_, err = io.Copy(fileWriter, file)
	if err != nil {
		panic(err)
	}

	contentType := bodyWriter.FormDataContentType()
	bodyWriter.Close()
	resp, err := http.Post(url, contentType, bodyBuf)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	log.Println(string(respBody))
}
