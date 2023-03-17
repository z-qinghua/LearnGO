// @program:     LearnGo
// @file:        main.go
// @create:      2023-03-15 22:12
// @description:

//用切片读写文件

//结合了flag包，和缓冲读写

package main

import (
    "flag"
    "fmt"
    "os"
)

func cat(f *os.File) {

    const NBUF = 512
    var buf [NBUF]byte
    for true {
        switch nr, err := f.Read(buf[:]); true {
        case nr < 0:
            fmt.Fprintf(os.Stderr, "cat:error reading:%s\n", err.Error())
            os.Exit(1)
        case nr == 0: //EOF
            return
        case nr > 0:
            if nw, ew := os.Stdout.Write(buf[0:nr]); nw != nr {
                fmt.Fprintf(os.Stderr, "cat:error writing:%s\n", ew.Error())
            }
        }
    }
}

func main() {
    flag.Parsed() //scans the arg list and sets up flags
    if flag.NArg() == 0 {
        cat(os.Stdin)
    }
    for i := 0; i < flag.NArg(); i++ {
        f, err := os.Open(flag.Arg(i))
        //Open打开一个文件用于读取。如果操作成功，返回的文件对象的方法可用于读取数据；对应的文件描述符具有O_RDONLY模式
        //open()只能用于读取文件
        if f == nil {
            fmt.Fprintf(os.Stderr, "cat:can't open %s err:%s\n", flag.Arg(i), err)

            //fmt.Fprintf() 依据指定的格式向第一个参数内写入字符串，第一参数必须实现了 io.Writer 接口。Fprintf() 能够写入任何类型，
            //只要其实现了 Write 方法，包括 os.Stdout, 文件（例如 os.File），管道，网络连接，通道等等，同样的也可以使用 bufio 包中缓冲写入。
            //bufio 包中定义了 type Writer struct{...}

            os.Exit(1)
        }
        cat(f)
        f.Close()
    }
}
