package main

import (
	"fmt"
)

/**

The definition of Higher-order function from wiki is a function which does at least one of the following

	1.takes one or more functions as arguments
	2.returns a function as its result
	以一个或多个函数作为参数
返回一个函数作为其结果
Let's look at some simple examples for the above two scenarios.
*/

//函数作为参数
func simple(a func(a, b int) int) {
	fmt.Println(a(60, 7))
}

func main() {
	//函数作为结果，赋值给变量
	f := func(a, b int) int {
		return a + b
	}
	simple(f)
}
