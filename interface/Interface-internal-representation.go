package main

import (
	"fmt"
)

type Describer interface {
	Describe()
}
type Person struct {
	name string
	age  int
}

func (p Person) Describe() { //implemented using value receiver
	fmt.Printf("%s is %d years old\n", p.name, p.age)
}

type Address struct {
	state   string
	country string
}

func (a *Address) Describe() { //implemented using pointer receiver
	//func (a Address) Describe() { //implemented using pointer receiver
	fmt.Printf("State %s Country %s", a.state, a.country)
}

func findType(i interface{}) {
	switch i.(type) {
	case string:
		fmt.Printf("I am a string and my value is %s\n", i.(string))
	case int:
		fmt.Printf("I am an int and my value is %d\n", i.(int))
	default:
		fmt.Printf("Unknown type\n")
	}
}

func assert(i interface{}) {
	//s := i.(int) //get the underlying int value from i
	//	fmt.Println(s)
	v, ok := i.(int)
	fmt.Println(v, ok)

}

//Empty Interface
func describe(i interface{}) {
	fmt.Printf("Type = %T, value = %v\n", i, i)
}

type Test interface {
	Tester()
}

type MyFloat float64

func (m MyFloat) Tester() {
	fmt.Println(m)
}

func describe10(t Test) {
	fmt.Printf("Interface type %T value %v\n", t, t)
}

func main101() {
	var t Test
	f := MyFloat(89.9)
	t = f
	describe10(t)
	t.Tester()

	s := "Hello World"
	describe(s)
	i := 55
	describe(i)
	strt := struct {
		name string
	}{
		name: "Naveen R",
	}
	describe(strt)

	var s1 interface{} = 56
	assert(s1)

	var s2 interface{} = "Steven Paul"
	assert(s2)

	findType("Naveen")
	findType(77)
	findType(89.98)

	var d1 Describer
	p1 := Person{"Sam", 25}
	d1 = p1
	d1.Describe()
	p2 := Person{"James", 32}
	d1 = &p2
	d1.Describe()

	var d2 Describer
	a := Address{"Washington", "USA"}

	/* compilation error if the following line is
	   uncommented
	   cannot use a (type Address) as type Describer
	   in assignment: Address does not implement
	   Describer (Describe method has pointer
	   receiver)
	*/
	//d2 = a

	d2 = &a //This works since Describer interface
	//is implemented by Address pointer in line 22
	d2.Describe()
}
