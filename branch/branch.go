// @program:     LearnGo
// @file:        branch.go
// @create:      2022-09-18 21:54
// @description:

package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

func grade(score int) string {
	g := ""
	switch {
	case score < 0 || score > 100:
		panic(fmt.Sprintf(
			"Wrong score:%d", score))
	case score < 60:
		g = "F"
	case score < 80:
		g = "C"
	case score < 90:
		g = "B"
	case score <= 100:
		g = "A"
	}
	return g
}

func convertToBin(n int) string {
	result := ""
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	return result

}

func printFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func main() {
	const filename = "abc.txt"
	//contents, err := ioutil.ReadFile(filename)
	//if err != nil {
	//	fmt.Println(err)
	//} else {
	//	fmt.Printf("%s\n", contents)
	//}
	if contents, err := ioutil.ReadFile(filename); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", contents)
	}
	fmt.Println(
		grade(0),
		grade(90),
		grade(100))

	fmt.Printf("--------------------------\n")
	fmt.Println(
		convertToBin(5),  //101
		convertToBin(13), //1101
		convertToBin(10)) //1010
	printFile("abc.txt")
}
