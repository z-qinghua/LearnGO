// @program:     LearnGo
// @file:        server.go
// @create:      2023-03-27 22:17
// @description: 
package main

import (
    "fmt"
    "net"
)

func main() {
    fmt.Println("starting the server:")
    //创建listener
    listener, err := net.Listen("tcp", "localhost:50000")
    if err != nil {
        fmt.Println("Error  listening", err.Error())
        return //终止程序
    }
    //监听客户端的连接
    for {
        conn, err := listener.Accept()
        if err != nil {
            fmt.Println("Error accepting", err.Error())
            return //终止程序
        }
        go doServerStuff(conn)
    }
}

func doServerStuff(conn net.Conn) {
    for {
        buff := make([]byte, 512)
        len, err := conn.Read(buff)
        if err != nil {
            fmt.Println("Error reading", err.Error())
            return //终止程序
        }
        fmt.Printf("received data:%v", string(buff[:len]))
    }
}
