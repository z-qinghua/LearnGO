// @program:     LearnGo
// @file:        middleware.go
// @create:      2023-03-19 20:23
// @description:

package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

//1 实现一个http server
//2 实现一个handler：hello
//3 实现中间件的功能：1）记录请求的url和方法 2）记录请求的网络的地址 3）记录方法的执行时间

func tracing(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("记录请求的url和方法: %s, %s", r.URL, r.Method)
		next.ServeHTTP(w, r)
	})
}

func logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("记录请求的网络的地址: %s", r.RemoteAddr)
		next.ServeHTTP(w, r)
	})
}

func timeRecording(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		next.ServeHTTP(w, r)
		elapsedTime := time.Since(startTime)

		log.Printf("记录方法的执行时间: %s", elapsedTime)

	})
}

func hello(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "hello")
}

func main() {
	http.Handle("/", timeRecording(logging(tracing(http.HandlerFunc(hello)))))
	http.ListenAndServe("localhost:8080", nil)
}
