package main

import "fmt"

func change(val *int) {
	*val = 55
}

func modify(arr *[3]int) {
	(*arr)[0] = 90
}

func modify1(arr *[3]int) {
	arr[0] = 90
}

func modify2(sls []int) {
	sls[0] = 90
}

func main() {
	b2 := 255
	var a2 *int = &b2
	fmt.Printf("Type of a is %T\n", a2)
	fmt.Println("address of b is", a2)

	a1 := 210
	var b1 *int
	if b1 == nil {
		fmt.Println("b is", b1)
		b1 = &a1
		fmt.Println("b after initialization is", b1)
	}

	b3 := 255
	a3 := &b3
	fmt.Println("address of b is", a3)
	fmt.Println("value of b is", *a3)
	*a3++
	fmt.Println("new value of b is", b3)

	a := 58
	fmt.Println("value of a before function call is", a)
	b := &a
	change(b)
	fmt.Println("value of a after function call is", a)

	a5 := [3]int{89, 90, 91}
	modify(&a5)
	fmt.Println(a5)

	a6 := [3]int{89, 90, 91}
	modify(&a6)
	fmt.Println(a6)

	a7 := [3]int{89, 90, 91}
	modify2(a7[:])
	fmt.Println(a7)
}
