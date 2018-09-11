package main

import (
	"fmt"
)

func printrMutlArray(ary [3][2]string) {
	for _, v1 := range ary {
		for _, v2 := range v1 {
			fmt.Printf("%s ", v2)
		}
		fmt.Printf("\n")
	}
}

func main() {
	len := 10
	for i := 0; i < len; i++ {
		fmt.Printf("%d ", i)
	}
	numsRange := []int{1, 10, 100, 1001}
	fmt.Println("")
	for num := range numsRange {
		fmt.Printf("%d->%d\n", num, numsRange[num])
	}
	/* for {
		fmt.Println("Hello World")
		time.Sleep(1)
	} */
	flag := 101
	switch flag {
	case 1, 11, 111:
		fmt.Printf("%s ", "t1")
	case 10:
		fmt.Printf("%s ", "t11")
	case 2:
		fmt.Printf("%s ", "t2")
	case 3:
		fmt.Printf("%s ", "t3")
	default:
		fmt.Printf("%s ", "default")

	}
	num := 75
	switch {
	case num >= 0 && num < 50:
		fmt.Println("0~50")
	case num >= 50 && num < 100:
		fmt.Println("50~100")
	}

	ary := [3]int{10}
	for _, v := range ary {
		fmt.Println(v)
	}

	ary1 := [...]float64{90.1, 90, 51, 45.9, 81.9}
	sum := float64(0)
	for i, v := range ary1 {
		fmt.Printf("index is %d element is %f\n", i, v)
		sum += v
	}
	fmt.Println(sum)
	multiAry := [3][2]string{
		{"lion", "triger"},
		{"cat", "dog"},
		{"col1", "col2"},
	}

	printrMutlArray(multiAry)

}
