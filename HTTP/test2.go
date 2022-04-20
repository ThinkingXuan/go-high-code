package main

import (
	"fmt"
	"net/http"
)

// 自定义net/http的handler

type Engine struct {
}

func (e *Engine) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/":
		fmt.Fprintf(rw, "URL.Path = %q\n", req.URL.Path)
	case "printhead":
		for k, v := range req.Header {
			fmt.Fprintf(rw, "Header[%q] = %q\n", k, v)
		}
	default:
		fmt.Fprintf(rw, "404 NOT FOUND: %s\n", req.URL)
	}

}

func main() {
	e := Engine{}
	http.ListenAndServe(":8080", &e)
}
