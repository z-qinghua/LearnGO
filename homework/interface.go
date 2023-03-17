// @program:     LearnGo
// @file:        interface.go
// @create:      2022-10-07 17:00
// @description:

package main

type Fisher interface {
	swim(int, int) (int, error)
}

type Animaler interface { //定义接口。通常接口名以er结尾
	move(int)
}

type Frog struct {
}

func (Frog) swim(int, int) (int, error) {
	return 1, nil
}

func (Frog) move(int) {

}

func main() {
	var t Fisher
	frog := Frog{}
	t = frog
	t.swim(1, 2)

	var h Animaler
	h = frog
	h.move(2)
}
