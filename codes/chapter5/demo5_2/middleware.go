package main

import (
	"io"
	"log"
	"net/http"
	"time"
)

// Step1: 实现一个HTTP Server
// Step2: 实现一个HTTP Handler
// Step3: 实现中间件的功能：
//	(1)、实现HTTP中间件记录请求的URL、方法。
//	(2)、实现HTTP中间件记录请求的网络地址。
//	(3)、实现HTTP中间件记录请求的耗时。

// 记录请求的URL和方法
func tracing(next http.HandlerFunc) http.HandlerFunc {
	return func(resp http.ResponseWriter, req *http.Request) {
		log.Printf("实现HTTP中间件记录请求的URL和方法：%s, %s", req.URL, req.Method)
		next.ServeHTTP(resp, req)
	}
}

// 记录请求的网络地址
func logging(next http.HandlerFunc) http.HandlerFunc {
	return func(resp http.ResponseWriter, req *http.Request) {
		log.Printf("实现HTTP中间件记录请求的网络地址：%s", req.RemoteAddr)
		next.ServeHTTP(resp, req)
	}
}

// 记录请求的耗时
func processing(next http.HandlerFunc) http.HandlerFunc {
	return func(resp http.ResponseWriter, req *http.Request) {
		start := time.Now()
		next.ServeHTTP(resp, req)
		duration := time.Since(start)
		log.Printf("实现HTTP中间件记录请求的耗时: %v", duration)
	}
}

func HelloHandler(resp http.ResponseWriter, req *http.Request) {
	io.WriteString(resp, "hello world")
}

func main() {
	http.HandleFunc("/", tracing(logging(processing(HelloHandler))))
	log.Printf("starting http server at: %s", "http://127.0.0.1:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
