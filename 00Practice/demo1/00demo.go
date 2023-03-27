package main

import "fmt"

func main() {
	// 1)统计3个班级成绩情况，每个班有5名同学
	// 求出各个班的平均分和所有班级的平均分[学生的成绩从键盘输入]

	// 分析实现思路1
	// 1. 统计1个班级情况，每个班级5名同学，求出该班级平均分
	// 2. 学生数量5个[先死后活]
	// 3. 定义一个变量sum统计班级的总分

	// 分析实现思路2
	// 1. 统计3个班级成绩情况，每个班有5名同学，求出每个班的平均分
	// 2. j代表第几个班级
	// 3. 定义一个变量存放总成绩

	// 分析实现思路3
	// 1. 把代码做活
	// 2. 定义两个变量，表示班级的个数和班级的人数

	// 分析思路
	// 1. 定义变量passNum统计及格人数

	var classNum int = 3
	var stuNum int = 5
	var totalSum float64 = 0.0
	var passNum int
	for j := 1; j <= classNum; j++ {
		var sum float64 = 0.0
		for i := 1; i <= stuNum; i++ {
			var score float64
			fmt.Printf("请输入第%d班第%d个学生的成绩\n", j, i)
			fmt.Scanln(&score)

			// 判断是否及格，并统计人数
			if score >= 60 {
				passNum++
			}

			// 累计总分
			sum = sum + score
		}
		fmt.Printf("第%d个班级平均分%v\n", j, sum/float64(stuNum))
		totalSum = totalSum + sum
	}
	fmt.Printf("各个班级的总成绩%v 所有班级平均分%v\n", totalSum, totalSum/(float64(classNum*stuNum)))
	fmt.Printf("共%d个学生及格", passNum)
}
