package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	// 创建一个 byte 类型的 26 个元素的数组，分别 放置'A'-'Z‘。
	// 使用 for 循环访问所有元素并打印出来。提示：字符数据运算 'A'+1 -> 'B'

	// 思路：
	// 1. 声明一个数组
	// 2. 使用for循环，利用字符可以进行运算的特点来赋值
	// 3. 使用for打印即可

	var mychars [26]byte
	for i := 0; i < 26; i++ {
		mychars[i] = 'A' + byte(i) // 注意需要将i => byte
	}
	for i := 0; i < 26; i++ {
		fmt.Printf("%c ", mychars[i])
	}

	// 请求出一个数组的最大值，并得到对应的下标
	// 1. 声明一个数组
	// 2. 假定第一个元素就是最大值，下标0
	// 3. 从第二个元素开始循环比较，如果发现有更大，则交换

	fmt.Println()
	var intArr = [...]int{1, -1, 9, 90, 11, 9000}
	maxVal := intArr[0]
	maxValIndex := 0
	for i := 1; i < len(intArr); i++ {
		if maxVal < intArr[i] {
			maxVal = intArr[i]
			maxValIndex = i
		}
	}
	fmt.Printf("maxVal=%v maxValIndex=%v\n", maxVal, maxValIndex)

	// 请求出一个数组的和和平均值。for-range
	// 思路：
	// 1. 声明一个数组
	// 2. 求出和sum
	// 3. 求出平均值

	var intArr2 = [...]int{1, -1, 9, 90, 11, 9000}
	sum := 0
	for _, val := range intArr2 {
		// 累计求和
		sum += val
	}
	// 让平均值保留到小数
	fmt.Printf("sum=%v 平均值=%v\n", sum, float64(sum)/float64(len(intArr)))

	// 要求：随机生成五个数，并将其反转打印 , 复杂应用.
	// 思路：
	// 1. 随机生成5个数，rand.Intn()函数
	// 2. 当我们得到随机数后，就放到一个数组
	// 3. 反转打印，交换次数len / 2，倒数第2个元素和第2个元素交换

	var intArr3 [5]int
	// 为了每次生成的随机数不一样，我们需要一个seed值
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < len(intArr3); i++ {
		intArr3[i] = rand.Intn(100) // 0<=n<=100
	}
	fmt.Println("交换前~=", intArr3)
	// 反转打印，交换的次数是len / 2
	// 倒数第一个和第一个元素交换，倒数第2个和第2个元素交换
	temp := 0	// 做一个临时变量
	for i := 0; i < len(intArr3) / 2; i++ {
		temp = intArr3[len(intArr3) - 1 - i]
		intArr3[len(intArr3) - 1 -i] = intArr3[i]
		intArr3[i] = temp
	}
	fmt.Println("交换后~=", intArr3)
}
