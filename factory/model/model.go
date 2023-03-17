package model

// type Student struct {
// 	Name  string
// 	Score float64
// }

// 如果 model 包的 结构体变量首字母小写，引入后，不能直接使用, 可以工厂模式解决
type student struct {
	Name  string
	score float64
}

// 一个工厂模式的函数，首字母大写。类似一个构造函数
func NewStudent(n string, s float64) *student { //返回的是结构体指针
	return &student{
		Name:  n,
		score: s,
	}
}

// 如果score字段首字母小写，在其他包不可以直接访问，我们可以提供一个方法
func (s *student) GetScore() float64 {
	return s.score
}

func (s *student) SetScore(score float64) {
	s.score = score
}
