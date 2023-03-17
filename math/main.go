// @program:     LearnGo
// @file:        main.go
// @create:      2022-10-17 21:07
// @description:

//数学常用函数
package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func constant() {
	fmt.Println(math.Pi)
	fmt.Println(math.E) //输出e

	//最大有符号整数
	maxInt := math.MaxInt
	fmt.Printf("%b\n", maxInt)

	//最大无符号整数
	umaxInt := uint64(math.MaxUint)
	fmt.Printf("%b\n", umaxInt)

	fmt.Println(math.SmallestNonzeroFloat32)
	fmt.Println(math.SmallestNonzeroFloat64)
}

func mathFunc() {
	//取整
	fmt.Println(math.Ceil(2.5))   //向上取整：3
	fmt.Println(math.Floor(2.5))  //向下取整：2
	fmt.Println(math.Ceil(-2.5))  //-2
	fmt.Println(math.Floor(-2.5)) //-3

	//截取整数部分
	fmt.Println(math.Trunc(2.5))  //2
	fmt.Println(math.Trunc(-2.5)) //-2

	//取绝对值
	fmt.Println(math.Abs(-3.6)) //3.6

	//分别取整数部分和小数部分
	fmt.Println(math.Modf(2.5))

	fmt.Println(math.Max(3, 7)) //返回大值
	fmt.Println(math.Min(3, 7)) //返回小值

	fmt.Println(math.Dim(3, 7)) //if x-y>0：x-y else: 0
	fmt.Println(math.Dim(7, 3))

	fmt.Println(math.Sqrt(9))        //开根号
	fmt.Println(math.Pow(3, 2))      //x的y次方
	fmt.Println(math.Pow10(3))       //10的3次方
	fmt.Println(math.Pow(math.E, 2)) //e的2次方
	fmt.Println(math.Exp(2))         //e的2次方
}

func exp_log() {
	fmt.Println(math.Log(math.E)) //默认底数e
	fmt.Println(math.Log1p(4))    //<=>math.Log(4 + 1)
	fmt.Println(math.Log(4 + 1))
	fmt.Println(math.Log2(4))     //2
	fmt.Println(math.Log10(1000)) //3
}

func trigo() {
	fmt.Println(math.Sin(1))
	fmt.Println(math.Cos(1))
	fmt.Println(math.Tanh(1)) //双矩阵
}

func ran() {
	rand.Seed(time.Now().UnixMilli()) //传入一个时间戳种子
	fmt.Println(rand.Int())           //生成一个随机整数
	fmt.Println(rand.Int31n(26))      //生成一个26以内的随机整数
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	rand.Shuffle(len(arr), func(i, j int) { //随机打乱
		arr[i], arr[j] = arr[j], arr[i]
	})
	fmt.Println(arr)

	//随机数生成器
	seed := time.Now().UnixMilli()
	source := rand.NewSource(seed)
	rander1 := rand.New(source)
	fmt.Println(
		rander1.Intn(26), rander1.Intn(26), rander1.Intn(26), rander1.Intn(26)) //输出随机生成的数

	source.Seed(seed) //重置种子，输出和上面一致
	rander2 := rand.New(source)
	fmt.Println(
		rander2.Intn(26), rander2.Intn(26), rander2.Intn(26), rander2.Intn(26)) //输出随机生成的数：输出和上面一致
}

func main() {
	constant()
	mathFunc()
	exp_log()
	trigo()
	ran()
}
