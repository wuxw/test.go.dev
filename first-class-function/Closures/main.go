package main

import (
	"fmt"
)

/*

Every closure is bound to its own surrounding variable.
每个闭包都与它自己的周围变量绑定在一起。让我们通过一个简单的例子来理解这意味着什么。
*/

func appendStr() func(string) string {
	t := "Hello"
	//闭包会把变量保持
	c := func(b string) string {
		t = t + " " + b
		return t
	}
	return c
}

/*

In the program above, the function appendStr returns a closure. This closure is bound to the variable t. Let's understand what this means.

The variables a and b declared in line nos. 16, 17 are closures and they are bound to their own value of t.

We first call a with the parameter World. Now the value of a's version of t becomes Hello World.

In line no. 20 we call b with the parameter Everyone. Since b is bound to its own variable t, b's version of t has a initial value of Hello again. Hence after this function call, the value of b's version of t becomes Hello Everyone. The rest of the program is self explanatory.

在上面的程序中，函数附属物返回一个闭包。这个闭包被绑定到变量t，让我们理解这是什么意思。
在第16、17行中声明的变量a和b是闭包，它们绑定到它们自己的t值。
我们首先调用一个参数世界。现在a的t版本的值变成了Hello World。
不一致。我们称b为参数每个人。因为b被绑定到它自己的变量t，所以b的t版本又有了Hello的初始值。因此在这个函数调用之后，b的t版本的值就变成了Hello。程序的其余部分是自解释的
*/
func main() {
	a1 := 5
	func() {
		a1 := 10
		fmt.Println("a =", a1)
	}()

	fmt.Println(a1)

	a := appendStr()
	b := appendStr()
	fmt.Println(a("World"))
	fmt.Println(b("Everyone"))

	fmt.Println(a("Gopher"))
	fmt.Println(b("!"))
}
