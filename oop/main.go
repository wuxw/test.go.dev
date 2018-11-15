package main

import "test.go.dev/oop/employee"

/*
We have made some important changes here. We have made the starting letter e of Employee struct to lower case, that is we have changed type Employee struct to type employee struct. By doing so we have successfully unexported the employee struct and prevented access from other packages. It's a good practice to make all fields of a unexported struct to be unexported too unless there is a specific need to export them. Since we don't need the fields of the employee struct anywhere outside the package, we have unexported all the fields too.

We have changed the field names accordingly in LeavesRemaining() method.

Now since employee is unexported, it's not possible to create values of type Employee from other packages. Hence we are providing a exported New function in line no. 14 which takes the required parameters as input and returns a newly created employee.

This program still has changes to be made to make it work but lets run this to understand the effect of the changes so far. If this program is run it will fail with the following compilation error,

go/src/constructor/main.go:6: undefined: employee.Employee
This is because we have unexported Employee and hence the compiler throws error that this type is not defined in main.go. Perfect. Just what we wanted. Now no other package will be able to create a zero valued employee. We have successfully prevented a unusable employee struct value from being created. The only way to create a employee now is to use the New function.


我们在这里做了一些重要的改变。我们已经将员工结构的起始字母e改为小写，也就是说我们已经更改了type Employee struct来输入Employee struct。通过这样做，我们成功地取消了员工结构，并阻止了其他软件包的访问。除非有特定的需求导出，否则将未导出的结构的所有字段都不导出也是一种很好的做法。因为我们不需要在包之外的任何地方的员工结构，所以我们也没有导出所有的字段。
*/
func main() {
	/* 	e := employee.Employee{
		FirstName:   "Sam",
		LastName:    "Adolf",
		TotalLeaves: 30,
		LeavesTaken: 20,
	} */
	//var e employee.Employee
	e := employee.New("Sam", "Adolf", 30, 20)
	e.LeavesRemaining()
}
