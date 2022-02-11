package main

import (
	"net/http"
)

func main() {
	// 注册路由
	http.HandleFunc("/", Hello)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
