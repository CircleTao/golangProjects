package main

import (
	"fmt"
	"log"
	"net/http"
)

// Engine is the uni handler for all requests
type Engine struct{}

// ResponseWriter 可以构造针对该请求的响应
// Request对象包含了该HTTP请求的所有的信息，比如请求地址、Header和Body等信息
func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/":
		fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
	case "/hello":
		for k, v := range req.Header {
			fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
		}
	default:
		fmt.Fprintf(w, "404 NOT FOUND: %s\n", req.URL)
	}
}

func main() {
	engine := new(Engine)
	// 传入engine实例，由于实现了ServeHTTP方法，使得所有的HTTP请求都转向我们自己的处理逻辑
	log.Fatal(http.ListenAndServe(":9999", engine))
}
