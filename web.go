package main

import (
	"fmt"
	"net/http"
)

func main() {
	// 创建一个HTTP服务器
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, World!")
	})

	// 创建一个HTTP请求处理程序
	http.ListenAndServe(":8080", nil)
}
