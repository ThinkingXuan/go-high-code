package main

import (
	"fmt"
	"log"
	"net/http"
)

// 测试net/http

func main() {
	http.HandleFunc("/", helloWord)
	http.HandleFunc("/printHead", printHead)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func helloWord(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("Hello World!")
}

func printHead(writer http.ResponseWriter, request *http.Request) {
	for key, value := range request.Header {
		fmt.Println(key, " ", value)
	}
}
