// @program:     LearnGo
// @file:        Map.go
// @create:      2022-09-20 21:24
// @description:

package main

import "fmt"

func main() {
	m := map[string]string{
		"name":    "ccmouse",
		"course":  "golang",
		"site":    "imooc",
		"quality": "notbad",
	}

	m2 := make(map[string]int) //m2 == Empty map

	var m3 map[string]int //m3 == nil

	fmt.Println(m, m2, m3)

	fmt.Println("Traversing map")
	for k, v := range m {
		fmt.Println(k, v)
	}

	fmt.Println("Getting values")
	courseName, ok := m["course"]
	fmt.Println(courseName, ok)
	if causeName, ok := m["cause"]; ok {
		fmt.Println(causeName)
	} else {
		fmt.Println("Key does not exit!")
	}

	fmt.Println("Deleting values")
	name, ok := m["name"]
	fmt.Println(name, ok)

	delete(m, "name")
	name, ok = m["name"]
	fmt.Println(name, ok)
}
