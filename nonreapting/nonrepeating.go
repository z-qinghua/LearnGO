// @program:     LearnGo
// @file:        nonrepeating.go
// @create:      2022-09-21 10:56
// @description:

//寻找最长不含有重复字符的子
/*
对于每个字母x
1-lastOccurred[x]不存在，或者< start ->无需操作
2-lastOccurred[x] >=start -> 更新start
3-更新lastOccurred[x]，更新maxlength
*/

package main

import (
	"fmt"
)

func lengthOfRepeatingSubStr(s string) int {
	lastOccurred := make(map[byte]int) // map
	start := 0
	maxlength := 0

	for i, ch := range []byte(s) {
		if lastI, ok := lastOccurred[ch]; ok && lastI >= start {
			start = lastI + 1
		}
		if i-start+1 > maxlength {
			maxlength = i - start + 1
		}
		lastOccurred[ch] = i
	}
	return maxlength
}

func main() {
	fmt.Println(lengthOfRepeatingSubStr("abcabcbb"))
	fmt.Println(lengthOfRepeatingSubStr("bbbbb"))
	fmt.Println(lengthOfRepeatingSubStr("pwwkew"))
}
