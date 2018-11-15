package employee

import (
	"fmt"
)

//type Employee struct {
/* type Employee struct {
	FirstName   string
	LastName    string
	TotalLeaves int
	LeavesTaken int
}
*/
type employee struct {
	firstName   string
	lastName    string
	totalLeaves int
	leavesTaken int
}

//构造函数

func New(firstName string, lastName string, totalLeave int, leavesTaken int) employee {
	e := employee{firstName, lastName, totalLeave, leavesTaken}
	return e
}

//func (e Employee) LeavesRemaining() {
func (e employee) LeavesRemaining() {
	//fmt.Printf("%s %s has %d leaves remaining", e.FirstName, e.LastName, (e.TotalLeaves - e.LeavesTaken))
	fmt.Printf("%s %s has %d leaves remaining", e.firstName, e.lastName, (e.totalLeaves - e.leavesTaken))
}
