package main

import (
	"fmt"
	"math"
	"sort"
)

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

type Person struct {
	Name string
	Age int
}

type ByAge []Person

func (list ByAge) Len() int {
	return len(list)
}

func (list ByAge) Swap (i, j int) {
	list[i], list[j] = 	list[j], list[i]
}

func (list ByAge) Less(i, j int) bool {
	return  list[i].Age < list [j].Age
}

func main() {
	d := []int{5, 2, 6, 3, 1, 4} // unsorted
	sort.Sort(sort.IntSlice(d))
	fmt.Println("float64 ", sort.SearchInts(d,6))
	fmt.Println(d)

	a := []float64{5.5, 2.2, 6.6, 3.3, 1.1, 4.4}
	sort.Sort(sort.Float64Slice(a))
	fmt.Println("float64 ", sort.SearchFloat64s(a,4.4))
	fmt.Println(a)
	// Output:[1.1 2.2 3.3 4.4 5.5 6.6]

	s := []string{"PHP", "golang", "python", "C", "Objective-C"}
	fmt.Println("float64 ", sort.SearchStrings(s,"golang"))
	sort.Sort(sort.StringSlice(s))
	fmt.Println(s)



	peoples := [] Person{
		{"bob" ,31},
		{"Joy" ,30},
		{"Michael" ,49},
		{"Jenny" ,81},
		{"Cat" ,24},
	}
	sort.Sort(ByAge(peoples))
	fmt.Println(peoples)
	dary := [...]int{1, 19, 81, 91, 1000, 1001, 80, 0, 67, 99}
	dslice1 := dary[:]
	//fmt.Println(SearchEle(dslice1, 67))
	fmt.Println("search ele ",SearchEle(dslice1, 0))
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


func BubbleSort(nums []int) {
/*	for i := 0; i < len(nums); i++ {
		for j := 1; j < len(nums)-i; j++ {
			if nums[j] < nums[j-1] {
				//交换
				nums[j], nums[j-1] = nums[j-1], nums[j]
			}
		}
	}*/
	length := len(nums)
	for i := 0; i < length ; i++ {
		for j := 1; j < length-1 ; j++ {
			if nums[j] < nums[j-1] {
				nums[j] ,nums[j-1] = nums[j-1] ,nums[j]
			}
		}
	}
}

func SearchEle (nums []int, ele int) int {
	isExist := -1
	 length := len(nums)
	 if length == 0 {
	 	return  isExist
	 }else if length == 1 {
		 if  ele == nums[0]  {
			return isExist
		 }
		 return  isExist
	 }
	BubbleSort(nums)
	// n :=  float64( length/2 )
	 l := int(math.Ceil(  float64( length/2 ) ))
	 if  ele == nums[l]  {
		isExist = l
	 }else if ele > nums[l] {
		for i := l; i < length ; i++ {
			if ele == nums[i] {
				isExist = i
			}
		}
	 }else{
		 for i := 0; i < l ; i++ {
			 if ele == nums[i] {
				 isExist = i
			 }
		 }
	 }
	 return isExist
}

/*
func bubbleSort(nums []int) {
	len := len(nums)
	for i := 0; i < len ; i++ {
	    for j := i+1; j < len ; j++ {
			if nums[i] < nums[j]{
				tmp := nums[i]
				nums[i] = nums[j]
				nums[j] = tmp
			}
	    }
	}
		
	}
}*/