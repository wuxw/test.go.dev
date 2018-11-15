package main

import (
	"fmt"
)

type student struct {
	firstName string
	lastName  string
	grade     string
	country   string
}

func filter(s []student, f func(student) bool) []student {
	var r []student
	for _, v := range s {
		if f(v) == true {
			r = append(r, v)
		}
	}
	return r
}

func iMap(s []int, f func(int) int) []int {
	var r []int
	for _, v := range s {
		r = append(r, f(v))
	}
	return r
}

/**
Here's a quick recap of what we learnt in this tutorial,

What are first class functions?
Anonymous functions
User defined function types
Higher-order functions
Passing functions as arguments to other functions
Returning functions from other functions
Closures
Practical use of first class functions
*/

func main() {
	//searchFirstName := "mike"
	/* hasFirstNameFn := func(s student) bool {
		if searchFirstName == s.firstName {
			return true
		}
		return false
	} */
	s1 := student{"mike", "white", "A", "usa"}
	s2 := student{
		firstName: "test",
		lastName:  "green",
		grade:     "Q",
		country:   "china",
	}

	var stuAry = []student{s1, s2}
	stuAry = filter(stuAry, func(s student) bool {
		//if searchFirstName == s.firstName {
		if "mike" == s.firstName {
			return true
		}
		return false
	})
	fmt.Println(stuAry)

	a := []int{5, 6, 7, 8, 9}
	r := iMap(a, func(n int) int {
		return n * 5
	})
	fmt.Println(r)
}
