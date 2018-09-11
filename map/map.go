package main

import "fmt"

func main() {
	var personSalary map[string]int
	if personSalary == nil {
		fmt.Println("map is nil. Going to make one.")
		personSalary = make(map[string]int)
	}

	personSalary["steve"] = 12000
	personSalary["jamie"] = 15000
	personSalary["mike"] = 9000
	fmt.Println("personSalary map contents:", personSalary)

	personSalary1 := map[string]int{
		"steve": 12000,
		"jamie": 15000,
	}
	personSalary1["mike"] = 9000
	fmt.Println("personSalary map contents:", personSalary1)

	employee := "jamie"
	fmt.Println("Salary of", employee, "is", personSalary[employee])
	fmt.Println("Salary of joe is", personSalary["joe"]) //return zero
	newEmp := "joe"
	value, ok := personSalary[newEmp] //这种赋值方法
	if ok == true {
		fmt.Println("Salary of", newEmp, "is", value)
	} else {
		fmt.Println(newEmp, "not found")
	}

	fmt.Println("delete before length is", len(personSalary))
	delete(personSalary, "steve")
	for k, v := range personSalary {
		fmt.Printf("personSalary1[%s] = %d\n", k, v)
	}
	fmt.Println("delete after length is", len(personSalary))
	newPersonSalary := personSalary
	newPersonSalary["mike"] = 18000
	fmt.Println("Person salary changed", personSalary)

}
