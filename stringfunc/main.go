package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	// 字符串遍历，同时解决中文乱码输出
	str := "hello北京"
	r := []rune(str)
	for i := 0; i < len(r); i++ {
		fmt.Printf("字符=%c\n", r[i])
	}

	// 字符串转整数
	n, err := strconv.Atoi("123")
	if err != nil {
		fmt.Println("转换错误")
	}
	fmt.Printf("n type=%T 整数：%d\n", n, n)

	// 整数转字符串
	str2 := strconv.Itoa(234)
	fmt.Printf("str2 type=%T 字符串：%q\n", str2, str2)

	// 字符串转byte
	var bytes = []byte("123")
	fmt.Printf("bytes=%c\n", bytes)

	// byte[]转字符串
	str3 := string([]byte{97, 98, 99})
	fmt.Printf("str3=%q\n", str3)

	// 十进制转对应2，8，16进制
	num1 := strconv.FormatInt(16, 2)
	fmt.Printf("对应的2进制=%v\n", num1)

	num2 := strconv.FormatInt(16, 8)
	fmt.Printf("对应的8进制=%v\n", num2)

	num3 := strconv.FormatInt(16, 16)
	fmt.Printf("对应的16进制=%v\n", num3)

	// 查找子串是否在指定的字符串中
	b := strings.Contains("seafood", "foo")
	fmt.Printf("b=%v\n", b)

	// 统计一个字符串中有几个指定的子串
	num := strings.Count("cechees", "e")
	fmt.Printf("num=%v\n", num)

	// 不区分字母大小写的字符串比较（==区分大小写）
	b2 := strings.EqualFold("abc", "abc")
	fmt.Printf("b=%v\n", b2)

	fmt.Println("结果=", "abc" == "ABC")

	// 返回子串在字符串中第一次出现的index值，如果没有则返回-1
	index := strings.Index("hello", "o")
	fmt.Printf("index=%v\n", index)

	// 返回字串在字符串中最后一次出现的index
	lastIndex := strings.LastIndex("gohello", "o")
	fmt.Printf("lastIndex=%v\n", lastIndex)

	// 将指定的子串替换成另外一个子串，并且可以指定替换几个，n=-1表示全部替换
	str2 = "go go hello"
	str = strings.Replace(str2, "go", "beijing", -1)
	fmt.Printf("str=%v str2=%v\n", str, str2) // 原本的字符串不会变

	// 按照某个指定的字符，为分割标识，将一个字符串拆分为字符串数组
	strArr := strings.Split("hello,world,ok", ",")
	for i := 0; i < len(strArr); i++ {
		fmt.Printf("strArr[%v]=%v\n", i, strArr[i])
	}
	fmt.Printf("strArr=%v\n", strArr)

	// 将字符串的字母进行大小写的转换
	str  = "golang Hello"
	str = strings.ToLower(str)
	fmt.Printf("str=%v\n",str) // golang hello
	str =strings.ToUpper(str) 
	fmt.Printf("str=%v\n",str) // GOLANG HELLO

	// 将字符串左右两边的空格去掉
	str = strings.TrimSpace(" tn a lone ntrm ")
	fmt.Printf("str=%q\n", str)

	// 将字符串左右两边指定的字符去掉
	str = strings.Trim(" ! hello!world,! ", " !")
	fmt.Printf("str=%q\n", str)

	// 将字符串左边指定的字符去掉
	str = strings.TrimLeft("hello world!", "he")
	fmt.Printf("str=%q\n", str)

	// 将字符串右边指定的字符去掉
	str = strings.TrimRight("hello,world", "world")
	fmt.Printf("str=%q\n", str)

	// 判断字符串是否以指定的字符串开头
	b = strings.HasPrefix("hello,world", "he")
	fmt.Printf("b=%v\n", b)

	// 判定字符串是否以指定的字符串结束
	b = strings.HasSuffix("hello,world", "world")
	fmt.Printf("b=%v\n", b)
}