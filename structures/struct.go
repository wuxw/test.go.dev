package main

import (
	"fmt"
	"model/computer"
)

type Employee1 struct {
	firstName string
	lastName  string
	age       int
}

type Employee struct {
	firstName, lastName string
	age, salary         int
}

//anonymous fields
type Person1 struct {
	string
	int
}

type Address struct {
	city, state string
}
type Person struct {
	name    string
	age     int
	address Address
}

//Promoted fields 提升 将city和state提升到1级，直接用.访问
type Person6 struct {
	name string
	age  int
	Address
}

func main() {
	emp1 := Employee{
		firstName: "Sam",
		lastName:  "Anderson",
		age:       30,
		salary:    1000,
	}
	emp2 := Employee{"Thomas", "Paul", 29, 9000}
	fmt.Println("Employee 1", emp1)
	fmt.Println("Employee 2", emp2)

	//独立使用，没有公共的 type Employee struct，所以是匿名结构体
	emp3 := struct {
		firstName, lastName string
		age, salary         int
	}{
		firstName: "Andreah",
		lastName:  "Nikola",
		age:       31,
		salary:    5000,
	}

	fmt.Println("Employee 3", emp3)

	var emp4 Employee //zero valued structure
	fmt.Println("Employee 4", emp4)

	emp6 := Employee{"Sam", "Anderson", 55, 6000}
	fmt.Println("First Name:", emp6.firstName)
	fmt.Println("Last Name:", emp6.lastName)
	fmt.Println("Age:", emp6.age)
	fmt.Printf("Salary: $%d", emp6.salary)

	emPoint8 := &Employee{"Sam", "Anderson", 55, 6000}
	fmt.Println("First Name:", (*emPoint8).firstName)
	fmt.Println("Age:", (*emPoint8).age)
	emp8 := *emPoint8
	fmt.Printf("%T\n", emPoint8)
	fmt.Println("Salary:", emp8.salary)

	var p1 Person1
	p1.string = "naveen"
	p1.int = 50
	fmt.Println(p1)

	var p Person
	p.name = "Naveen"
	p.age = 50
	p.address = Address{
		city:  "Chicago",
		state: "Illinois",
	}
	fmt.Println("Name:", p.name)
	fmt.Println("Age:", p.age)
	fmt.Println("City:", p.address.city)
	fmt.Println("State:", p.address.state)

	var p6 Person6
	p6.name = "Naveen"
	p6.age = 50
	p6.Address = Address{
		city:  "Chicago",
		state: "Illinois",
	}
	fmt.Println("Name:", p6.name)
	fmt.Println("Age:", p6.age)
	fmt.Println("City:", p6.city)   //city is promoted field
	fmt.Println("State:", p6.state) //state is promoted field

	var spec computer.Spec
	spec.Maker = "apple"
	spec.Price = 50000
	fmt.Println("Spec:", spec)

}
