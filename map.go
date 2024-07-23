package main

import "fmt"

func main() {
	m := map[int]int{1: 10, 2: 20, 3: 30}
	fmt.Println(m)

	// 判断键值存在
	if v,ok := m[3]; ok{
		fmt.Println(v)
	}

}