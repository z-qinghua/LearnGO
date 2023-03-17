// @program:     LearnGo
// @file:        main.go
// @create:      2022-10-07 13:49
// @description:

package main

type Transporter interface {
	move(string, string) (int, error)
	say(int)
}

type Humna interface {
	think(int)
}
type Car struct { //定义结构体时无需显式声明它要实现什么接口
}

func (Car) move(string, string) (int, error) {
	return 1, nil
}

func (Car) say(a int) {

}

func (*Car) think(a int) { //方法接收者用指针，则实现接口的是指针类型

}

func foot(t Transporter) {

}
func main() {
	var t Transporter
	car := Car{}
	t = car  //t并不是和Car等价
	t = &car //值实现的方法，指针同样也实现了
	t.say(2)

	foot(t)
	foot(car) //car虽然不是接口，但实现了接口

	var h Humna
	h = &car
	h.think(4)
	car.say(4)
	car.think(3)
}
