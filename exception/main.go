// @program:     LearnGo
// @file:        main.go
// @create:      2022-10-07 12:56
// @description:

package main

import (
	"errors"
	"fmt"
	"time"
)

type PathError struct {
	path      string
	op        string
	creatTime string
	message   string
}

func NewPathError(path, op, message string) PathError {
	return PathError{
		path:      path,
		op:        op,
		creatTime: time.Now().Format("2006-01-02"),
		message:   message,
	}
}

func (e PathError) Error() string { //实现error接口
	return e.creatTime + ":" + e.op + " " + e.path + " " + e.message
}

func deletePath(path string) error {
	//
	return NewPathError(path, "delete", "path not exits")
}

func soo() {
	defer func() {
		if err := recover(); err != nil { //1-recover会阻断panic的执行，2-recover必须在defer中才能生效
			fmt.Println("发生了panic，不让程序退出")
		}
	}()
	panic(errors.New("这里是错误信息"))

}

func main() {
	soo()
	fmt.Println("333333")
}
