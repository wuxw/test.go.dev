package main

import "fmt"

func subtactOne(numbers []int) {
	for i := range numbers {
		numbers[i] -= 2
	}

}

func countries() []string {
	countries := []string{"USA", "Singapore", "Germany", "India", "Australia"}
	neededCountries := countries[:len(countries)-2]
	//copy前，产生一个对应的新数组，预装
	countriesCpy := make([]string, len(neededCountries))
	fmt.Println("array before", countries)

	copy(countriesCpy, neededCountries) //copies neededCountries to countriesCpy
	for i, _ := range countriesCpy {
		countriesCpy[i] += "New"
	}
	fmt.Println("array before", countries)
	return countriesCpy
}

func copyArray(copyArray []string) []string {
	cpyAry := make([]string, len(copyArray))

	copy(cpyAry, copyArray)
	return cpyAry
}

func main() {
	dary := [...]int{1, 19, 81, 91, 1000, 1001, 80, 0, 67, 99}
	iStart := 3
	dslice := dary[iStart:7]
	fmt.Println("array before", dary)
	for i, _ := range dslice {
		dslice[i]++
	}
	fmt.Println("array after", dary)

	numa := [3]int{78, 79, 80}
	nums1 := numa[:] //creates a slice which contains all elements of the array
	nums2 := numa[:]
	fmt.Println("array before change 1", numa)
	nums1[0] = 100
	fmt.Println("array after modification to slice nums1", numa)
	nums2[1] = 101
	fmt.Println("array after modification to slice nums2", numa)

	cars := []string{"Ferrari", "Honda", "Ford"}
	fmt.Println("cars:", cars, "has old length", len(cars), "and capacity", cap(cars)) //capacity of cars is 3
	cars = append(cars, "Toyota")
	fmt.Println("cars:", cars, "has new length", len(cars), "and capacity", cap(cars)) //capacity of cars is doubled to 6

	nos := []int{8, 7, 6}
	fmt.Println("slice before function call", nos)
	//引用传递
	subtactOne(nos)                               //function modifies the slice
	fmt.Println("slice after function call", nos) //modifications are visible outside

	/*数组copy*/
	countriesNeeded := countries()
	fmt.Println(countriesNeeded)

}
