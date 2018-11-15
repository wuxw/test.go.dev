/* package main

import (
	"fmt"
)

func a() {
	n := []int{5, 7, 4}
	fmt.Println(n[3])
	fmt.Println("normally returned from a")
}
func main() {
	a()
	fmt.Println("normally returned from main")
}
*/

package main

import (
	"fmt"
)

func r() {
	if r := recover(); r != nil {
		fmt.Println("Recovered", r)
	}
}

func a() {
	defer r()
	n := []int{5, 7, 4}
	fmt.Println(n[3])
	fmt.Println("normally returned from a")
}

func main() {
	a()
	fmt.Println("normally returned from main")
}
