package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// 设置两个路由，分别绑定 indexHandler 和 helloHandler
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/hello", helloHandler)

	// 启动web服务，:9999表示在9999端口监听，nil 代表使用标准库中的实例处理
	log.Fatal(http.ListenAndServe(":9999", nil))
}

// handler echoes r.URL.Path
func indexHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
}

// handler echoes r.URL.Header
func helloHandler(w http.ResponseWriter, req *http.Request) {
	for k, v := range req.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
}
