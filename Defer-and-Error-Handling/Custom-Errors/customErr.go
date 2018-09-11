package main

import (
	"fmt"
	"math"
)

/**

Here is a quick recap of what we learnt in this tutorial,

	Creating custom errors using the New function
	Adding more information to the error using Errorf
	Providing more information about the error using struct type and fields
	Providing more information about the error using methods on struct types
	使用新函数创建自定义错误
	使用Errorf向错误添加更多信息
	使用struct类型和字段提供关于错误的更多信息
	使用struct类型的方法提供关于错误的更多信息
*/
func circleArea(radius float64) (float64, error) {
	if radius < 0 {
		//	return 0, errors.New("Area calculation failed, radius is less than zero")

		//Adding more information to the error using Errorf
		return 0, fmt.Errorf("Area calculation failed, radius %0.2f is less than zero", radius)
	}
	return math.Pi * radius * radius, nil
}

func main() {
	radius := -20.0
	area, err := circleArea(radius)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Area of circle %0.2f", area)
}
