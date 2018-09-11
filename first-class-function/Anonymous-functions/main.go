package main

import (
	"fmt"
)

//These kind of functions are called anonymous functions since they do not have a name.

func main() {

	//Anonymous functions
	a := func() {
		fmt.Println("hello world first class function")
	}
	a()
	fmt.Printf("%T", a)

	//Anonymous functions 自执行
	func() {
		fmt.Println("hello world first class function")
	}()

	//支持带参数
	func(n string) {
		fmt.Println("Welcome", n)
	}("Gophers")

}
