// @program:     LearnGo
// @file:        io.go
// @create:      2022-10-19 19:51
// @description:

// io操作
package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func std() {
	fmt.Println("please input 2 int")
	var s int
	fmt.Scanf("%d", &s)
	var a int
	fmt.Scanf("%d", &a)
	fmt.Printf("You input %d %d\n", s, a)
}

//读文件数据
func readFile1() {
	if file, err := os.Open("go.mod"); err != nil {
		fmt.Println(err) //假如文件不存在或者其它错误时会报错
	} else {
		defer file.Close()
		var buffer strings.Builder
		for {
			bs := make([]byte, 10)
			if n, err := file.Read(bs); err != nil {
				fmt.Println(err)
				if err == io.EOF { //读到文件末尾
					break
				}
			} else {
				fmt.Printf("从文件中读出%d个字节\n", n)
				buffer.WriteString(string(bs))
			}
		}
		fmt.Println(buffer.String())
	}
}

func readFile2() {
	if file, err := os.Open("go.mod"); err != nil {
		fmt.Println(err) //假如文件不存在或者其它错误时会报错
	} else {
		defer file.Close()
		reader := bufio.NewReader(file)
		var i int
		for {
			if line, err := reader.ReadString('\n'); err != nil {
				fmt.Println(err)
				if err == io.EOF { //读到文件末尾
					break
				}
			} else {
				//fmt.Print(line)
				fmt.Printf("第%d行\n", i)
				fmt.Println(strings.Trim(line, "\n"))
				i++
			}
		}
	}
}

//写文件数据
func writeFile1() {
	if file, err := os.OpenFile(
		"writeTest.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 06666); err != nil {
		fmt.Println(err)
	} else {
		defer file.Close()
		file.Write([]byte("helloWorld\n"))
	}
}

func writeFile2() { //效率更高
	if file, err := os.OpenFile(
		"writeTest.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 06666); err != nil {
		fmt.Println(err)
	} else {
		defer file.Close()
		writer := bufio.NewWriter(file)
		writer.WriteString("hello") //只把内容写到内存缓冲区
		writer.WriteString("\n")
		writer.Flush() //强行把缓冲区内容刷到磁盘里
	}
}

//文件的一些操作
func fileOp() {
	//os.Remove("writeTest.txt")  //删除文件
	//os.MkdirAll("p1/p2", 06666) //创建多级目录
	//os.Create("p1/p2/lala.txt") //创建文件
	//os.RemoveAll("p1")          //删除目录
	//os.Create("old.txt")
	//os.Rename("old.txt", "new.txt") //重新命名

	file, _ := os.Open("go.mod")
	fmt.Println(file.Name())
	info, _ := file.Stat()
	fmt.Println(info.Name())
	fmt.Println(info.Mode()) //权限
	fmt.Println(info.IsDir())
	fmt.Println(info.Size())    //文件大小
	fmt.Println(info.ModTime()) //最后一次修改时间
}

//遍历目录里目录及文件名字
func walkDir(path string) error {
	if fileInfos, err := ioutil.ReadDir(path); err != nil {
		return err
	} else {
		for _, fileInfo := range fileInfos {
			fmt.Println(fileInfo.Name()) //输出名字
			if fileInfo.IsDir() {
				if err := walkDir(filepath.Join(path, fileInfo.Name())); err != nil { //递归实现
					return err
				}
			}
		}
	}
	return nil
}

func logger() {
	if file, err := os.OpenFile(
		"mage.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666); err != nil {
		fmt.Println(err)
	} else {
		defer file.Close()
		logWriter := log.New(file, "[BIZ_NAME]", log.Ldate|log.Lmicroseconds) //自带的log实现
		logWriter.Println("hello")
		logWriter.Println("How are you")
	}
}

//调用系统命令
func sysCmd() {
	cmd_path, err := exec.LookPath("netstat") //which of
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(cmd_path)
		cmd := exec.Command("netstat", "-s") //文本不能有空格
		output, err := cmd.Output()
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(string(output))
		}
	}
	//cmd := exec.Command("del", "abc.txt")
	//if err := cmd.Run(); err != nil { //没有输出则直接运行
	//	fmt.Println(err)
	//}
}
func main() {
	//std()
	//readFile1()
	//readFile2()
	//writeFile1()
	//writeFile2()
	//fileOp()
	//walkDir(
	//	"E:\\我的文件\\personal files\\Study\\Code\\Golang\\workspace\\LearnGo\\src\\main") //完整路径或者相对项目路径
	//logger()
	sysCmd()
}
