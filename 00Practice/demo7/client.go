// @program:     LearnGo
// @file:        client.go
// @create:      2023-03-27 22:18
// @description: 

package main

import (
    "fmt"
    "os"
    "net"
    "bufio"
    "strings"
)

func main() {
    //打开连接
    conn, err := net.Dial("tcp", "localhost:50000")
    if err != nil {
        //由于目标计算机积极拒绝而无法创建连接
        fmt.Println("Error dialing", err.Error())
        return //终止程序
    }
    inputReader := bufio.NewReader(os.Stdin)
    fmt.Println("First, what is your name?")
    clientName, _ := inputReader.ReadString('\n')
    //fmt.Println("clientName:", clientName)

    trimmedClient := strings.Trim(clientName, "\r\n")
    //给服务器发送信息直到程序退出
    for {
        fmt.Println("what to send to the server?Type Q to quit.")
        input, _ := inputReader.ReadString('\n')
        trimmedInput := strings.Trim(input, "\r\n")
        //fmt.Println("input:", input)
        //fmt.Println("trimmedInput:", trimmedInput)

        if trimmedInput == "Q" {
            return
        }
        _, err := conn.Write([]byte(trimmedClient + " says:" + trimmedInput))
        if err != nil {
            fmt.Println("Error writing", err.Error())
        }
    }
}
