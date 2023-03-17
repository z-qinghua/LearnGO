package main

import (
	"LearnGo/0projectExcercise/familyAccount/utils"
	"fmt"
)

func main() {
	fmt.Println("用面向对象方式完成")
	utils.NewFamilyAccount().Menu()
}
