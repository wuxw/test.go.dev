package main

import (
	"fmt"
	"math"
)

type myInt int

func (a myInt) add(b myInt) myInt {
	return a + b
}

type rectangle1 struct {
	length int
	width  int
}

func perimeter(r *rectangle1) {
	fmt.Println("perimeter function output:", 2*(r.length+r.width))

}

func (r *rectangle1) perimeter() {
	fmt.Println("perimeter method output:", 2*(r.length+r.width))
}

type address struct {
	city  string
	state string
}

func (a address) fullAddress() {
	fmt.Printf("Full address: %s, %s", a.city, a.state)
}

type person struct {
	firstName string
	lastName  string
	address
}

type Employee struct {
	name     string
	salary   int
	currency string
	age      int
}

/*
 displaySalary() method has Employee as the receiver type
*/
func (e Employee) displaySalary() {
	fmt.Printf("Salary of %s is %s%d", e.name, e.currency, e.salary)
}

/*
Method with value receiver
*/
func (e Employee) changeName(newName string) {
	e.name = newName
}

/*
Method with pointer receiver
*/
func (e *Employee) changeAge(newAge int) {
	e.age = newAge
}

type Rectangle struct {
	length int
	width  int
}

type Circle struct {
	radius float64
}

func Area(r Rectangle) {
	fmt.Printf("Area Function result: %d\n", (r.length * r.width))
}

func (r Rectangle) Area() int {
	return r.length * r.width
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func main() {
	emp1 := Employee{
		name:     "Sam Adolf",
		salary:   5000,
		currency: "$",
		age:      33,
	}

	fmt.Printf("Employee name before change: %s", emp1.name)
	emp1.changeName("Michael Andrew")
	fmt.Printf("\nEmployee name after change: %s So no change", emp1.name)

	fmt.Printf("\n\nEmployee age before change: %d", emp1.age)
	(&emp1).changeAge(51)
	fmt.Printf("\nEmployee age after change: %d", emp1.age)
	fmt.Println("")
	emp1.displaySalary() //Calling displaySalary() method of Employee type

	fmt.Println("")

	r := Rectangle{
		length: 10,
		width:  5,
	}
	fmt.Printf("Area of rectangle %d\n", r.Area())
	c := Circle{
		radius: 12,
	}
	fmt.Printf("Area of circle %f", c.Area())

	fmt.Println("")

	p := person{
		firstName: "Elon",
		lastName:  "Musk",
		address: address{
			city:  "Los Angeles",
			state: "California",
		},
	}

	p.fullAddress() //accessing fullAddress method of address struct

	fmt.Println("")
	r1 := Rectangle{
		length: 10,
		width:  5,
	}
	Area(r1)
	r1.Area()

	p1 := &r1
	/*
	   compilation error, cannot use p (type *rectangle) as type rectangle
	   in argument to area
	*/
	//area(p)

	p1.Area() //calling value receiver with a pointer

	fmt.Println("")
	r9 := rectangle1{
		length: 10,
		width:  5,
	}
	p9 := &r9 //pointer to r
	perimeter(p9)
	p9.perimeter()

	/*
	   cannot use r (type rectangle) as type *rectangle in argument to perimeter
	*/
	//perimeter(r9)

	// will be interpreted by the language as (&r).perimeter()
	r9.perimeter() //calling pointer receiver with a value

	num1 := myInt(5)
	num2 := myInt(10)
	sum := num1.add(num2)
	fmt.Println("Sum is", sum)
}
