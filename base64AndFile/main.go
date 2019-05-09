package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("start----------------")
	fileToBase64()
	//base64ToFile()
	fmt.Println("end---------------")
}

func base64ToFile() {
	data, err := ioutil.ReadFile("test.txt")
	if err != nil {
		panic(err)
	}
	decodeData, err := base64.StdEncoding.DecodeString(string(data))
	if err != nil {
		panic(err)
	}
	f, err := os.OpenFile("test.pdf", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	f.Write(decodeData)
}

func fileToBase64() {
	data, err := ioutil.ReadFile("test.pdf")
	if err != nil {
		panic(err)
	}
	base64Str := base64.StdEncoding.EncodeToString(data)
	f, err := os.OpenFile("test.txt", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	f.Write([]byte(base64Str))
}
