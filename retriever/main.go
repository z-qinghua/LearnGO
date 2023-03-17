// @program:     LearnGo
// @file:        main.go
// @create:      2022-09-27 14:05
// @description:

package main

import (
	"LearnGo/src/main/retriever/mock"
	"LearnGo/src/main/retriever/real"
	"fmt"
)

type retriever interface {
	Get(url string) string
}

func download(r retriever) string {
	return r.Get("https://www.imooc.com")
}

func main() {

	var r retriever
	r = mock.Retriever{"this is a fake imooc.com"}
	r = real.Retriever{}
	fmt.Println(download(r))
}
