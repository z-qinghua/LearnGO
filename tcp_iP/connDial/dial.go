// @program:     LearnGo
// @file:        dial.go
// @create:      2023-03-27 22:34
// @description: 

//make a connection with www.example.org
package main

import (
    "net"
    "fmt"
    "os"
)

func main() {
    conn, err := net.Dial("tcp",
        "192.0.32.10:80") //tcp  ipv4
    checkConnection(conn, err)

    conn, err = net.Dial("udp", "192.0.32.10:80") //udp
    checkConnection(conn, err)

    conn, err = net.Dial("tcp", "[2620:0:2d0:200::10]:80") //ipv6
    checkConnection(conn, err)
}

func checkConnection(conn net.Conn, err error) {
    if err != nil {
        fmt.Printf("Error connecting %v\n", err.Error())
        os.Exit(1)
    }
    fmt.Printf("Connection is made with %v\n", conn)
}
