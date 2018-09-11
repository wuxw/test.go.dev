package main

import (
	"fmt"
)

//声明函数-结构体
type add func(a int, b int) int

func main() {

	//声明一个是add类型的函数结构体变量 addFn
	var addFn add = func(a int, b int) int {
		return a + b
	}
	s := addFn(5, 6)
	fmt.Println("Sum", s)
}
