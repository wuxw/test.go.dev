package main

import (
	"fmt"
	"reflect"
)

type order struct {
	ordId      int
	customerId int
}

func createQuery(q interface{}) {
	t := reflect.TypeOf(q)
	k := t.Kind()
	fmt.Println("Type ", t)
	fmt.Println("Kind ", k)

	if reflect.ValueOf(q).Kind() == reflect.Struct {
		v := reflect.ValueOf(q)
		t1 := reflect.TypeOf(q).Name()
		t := reflect.TypeOf(q)
		fmt.Println("T is ", t1)
		fmt.Println("Number of fields", v.NumField())

		fmt.Println("v is ", v)
		fieldNum := v.NumField()
		for i := 0; i < fieldNum; i++ {

			k := v.Field(i).Kind()
			fmt.Println("Kind is ", k)

			fmt.Println("v1 is ", v.Field(i).Int())
			fmt.Println("field is ", t.Field(i).Name)

			fmt.Printf("Field:%d type:%T value:%v\n", i, v.Field(i), v.Field(i))
		}
	}

}
func main() {
	o := order{
		ordId:      456,
		customerId: 56,
	}
	createQuery(o)

	a := 56
	x := reflect.ValueOf(a).Int()
	fmt.Printf("type:%T value:%v\n", x, x)
	b := "Naveen"
	y := reflect.ValueOf(b).String()
	fmt.Printf("type:%T value:%v\n", y, y)

}
