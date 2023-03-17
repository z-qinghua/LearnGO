// @program:     LearnGo
// @file:        byteAndruneAndstring.go
// @create:      2022-09-21 22:07
// @description:

package main

import "fmt"

//将字符串转换为rune的数组，并替换数组第i个元素为ch
func strToRune(str *string, i int, ch rune) {
	temp := []rune(*str)
	temp[i] = ch
	*str = string(temp)
}

//将字符串更换为byte的数组，并替换数组第i个元素为ch
func strToByte(str *string, i int, ch byte) {
	temp := []byte(*str)
	temp[i] = ch
	*str = string(temp)
}

func main() {
	str := "你好 hello"
	str1 := "你好 hello"
	str2 := "你好 hello"
	strToRune(&str, 0, 'A')  //将第0个字符替换为A
	strToByte(&str1, 0, 'A') //将第0个字符替换为A
	strToByte(&str2, 7, 'A') //将第7个字符替换为A
	fmt.Println("rune数组处理中文后：", str)
	fmt.Println("byte数组处理中文后：", str1)
	//转换成byte数组替换掉第0个字符，会出现地2、3字节无法显示问题
	fmt.Println("byte数组处理中文后：", str2)
}
