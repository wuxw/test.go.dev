package main

import (
	"fmt"
	"io/ioutil"
	"time"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func numbers() {
	for i := 1; i <= 5; i++ {
		time.Sleep(250 * time.Millisecond)
		fmt.Printf("%d ", i)
	}
}

func alphabets() {
	for i := 'a'; i <= 'e'; i++ {
		time.Sleep(400 * time.Millisecond)
		fmt.Printf("%c ", i)
	}
}

func hello() {
	wireteString := "Hello world goroutine"
	fileName := "D:/go-pro/src/test.go.dev/concurrency/goroutines/run.log"
	var d1 = []byte(wireteString)
	err2 := ioutil.WriteFile(fileName, d1, 0777) //写入文件(字节数组)
	check(err2)

	fmt.Println(wireteString)
}

func main() {
	go hello()
	//	time.Sleep(1 * time.Second)
	fmt.Println("main function")

	go numbers()
	go alphabets()
	time.Sleep(3000 * time.Millisecond)
	fmt.Println("main terminated")

}
