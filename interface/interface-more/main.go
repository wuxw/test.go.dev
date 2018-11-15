package main

import "fmt"

type I interface {
	Get() int

	Set(int)
}

type S struct {
	Age int
}

func (s S) Get() int {
	return s.Age
}
func (s *S) Set(age int) {
	s.Age = age
}

//2、interface 变量存储的是实现者的值
func f(i I) {
	//3、如何判断 interface 变量存储的是哪种类型
	//value, ok := em.(T)：em 是 interface 类型的变量，T代表要断言的类型，value 是 interface 变量存储的值，ok 是 bool 类型表示是否为该断言的类型 T。

	//如果需要区分多种类型，可以使用 switch 断言，更简单直接，这种断言方式只能在 switch 语句中使用。
	/* 	switch t := i.(type) {
	   	case *S:
	   		fmt.Println("i store *S", t)
	   	case *R:
	   		fmt.Println("i store *R", t)
	   	} */

	if t, ok := i.(*S); ok {
		fmt.Println("s implements I", t)
		i.Set(10)
		fmt.Println(i.Get())
		return
	}
	fmt.Println("s is not implements I")

}

func main() {
	//1、interface 是一种类型
	s := S{5}
	f(&s) //10

}
