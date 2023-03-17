package main

import (
	"fmt"
	"os"
)

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func main() {
	isPathExist, err := PathExists("E:/WorkSpace/goProject/src/LearnGo/file/write/abc.txt")
	if isPathExist {
		fmt.Println("Path exist!")
	}
	if !isPathExist {
		if err == nil {
			fmt.Println("Path do not exist!")
		} else {
			fmt.Printf("Path do not exist and err = %v\n", err)
		}
	}
}
