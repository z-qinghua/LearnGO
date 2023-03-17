// @program:     LearnGo
// @file:        ttt.go
// @create:      2022-10-12 15:54
// @description:

package main

import (
	"fmt"
	"time"
)

func basic() {

	//计算时间差
	begin := time.Now()
	fmt.Println(begin) //当前时间
	fmt.Println("fdsvr")

	//时间戳
	t := time.Now()
	fmt.Println(t.Unix())      //秒
	fmt.Println(t.UnixMilli()) //毫秒
	fmt.Println(t.UnixMicro()) //微秒
	fmt.Println(t.UnixNano())  //纳秒

	//方法一：
	useTime := time.Since(begin)

	//方法二：
	//end := time.Now()
	//useTime := end.Sub(begin)

	//输出时间差
	fmt.Println(useTime.Seconds())
	fmt.Println(useTime.Nanoseconds())

	//输出八小时后的时间
	dua := time.Duration(8 * time.Hour) //时间段：8小时
	end := begin.Add(dua)               //加上8小时
	fmt.Println(end)

	//Time 时刻
	//Duration  时间段
	//时刻 + 时间段 = 另一个时间
	//时刻 - 时刻 = 时间段
	// 不能： 时刻 + 时刻
	fmt.Println("-------------------------------")
	fmt.Println(begin.Year())       //年
	fmt.Println(begin.Month())      //月
	fmt.Println(int(begin.Month())) //月份数字表示
	fmt.Println(begin.Day())        //每月第几天
	fmt.Println(begin.YearDay())    //一年中第几天
	fmt.Println(begin.Weekday())    //周几
	fmt.Println(begin.Hour())       //时
	fmt.Println(begin.Minute())     //分
	fmt.Println(begin.Second())     //秒

}

const (
	TIME_FMT  = "2006-01-02 15:04:05"
	TIME_FMT2 = "2006-01-02"
	TIME_FMT3 = "20060102"
)

//时间格式化输出
func timeFmt() {
	//begin := time.Now()
	//fmt.Println(begin)
	//fmt.Println(begin.Format("2006-01-02 15:04:05"))
	fmt.Println(time.Now().Format(TIME_FMT))
	fmt.Println(time.Now().Format(TIME_FMT2))
	fmt.Println(time.Now().Format(TIME_FMT3))

	//字符串转化为时间格式

	//方法一：
	//if t, err := time.Parse(TIME_FMT2, "1999-10-15"); err == nil { //工作时不要这样写
	//	fmt.Println(t.Year())
	//	fmt.Println(t.Month())
	//	fmt.Println(t.Day())
	//}

	//方法二：
	loc, _ := time.LoadLocation("Asia/Shanghai")
	if t, err := time.ParseInLocation(TIME_FMT2, "1999-10-15", loc); err == nil { //一定要用ParseInLocation
		fmt.Println(t.Year())
		fmt.Println(t.Month())
		fmt.Println(int(t.Month())) //月份数字显示
		fmt.Println(t.Day())
	}
}

func ticker() {
	tc := time.NewTicker(3 * time.Second) //间隔3秒
	defer tc.Stop()
	for i := 0; i < 6; i++ {
		<-tc.C //阻塞3秒
		fmt.Println(time.Now().Unix())
	}
}

func timer() {
	fmt.Println(time.Now().Unix())
	tm := time.NewTimer(2 * time.Second)
	defer tm.Stop()
	<-tm.C //阻塞2秒钟
	fmt.Println(time.Now().Unix())

	for i := 0; i < 3; i++ {
		tm.Reset(1 * time.Second)
		<-tm.C
		fmt.Println(time.Now().Unix())
	}

	//<-time.After(4 * time.Second) //阻塞4秒钟，且不用stop
	//fmt.Println(time.Now().Unix())

	//time.Sleep(3 * time.Second)//阻塞3秒钟
	//fmt.Println(time.Now().Unix())
}
func main() {
	basic()
	timeFmt()
	ticker()
	timer()
}
