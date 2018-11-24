package main

import (
	"fmt"
)

type People interface {
	Speak(string) string
}

type Stduent struct{}

type People1 struct{}

func (p *People1) ShowA() {
	fmt.Println("people1 showA")
	p.ShowB()
}
func (p *People1) ShowB() {
	fmt.Println("people1 showB")
}


type Teacher struct {
	People1
}

func (t *Teacher) ShowB() {
	fmt.Println("teacher showB")
}


func (stu *Stduent) Speak(think string) (talk string) {
	if think == "bitch" {
		talk = "You are a good boy"
	} else {
		talk = "hi"
	}
	return
}

func main() {
	//编译不通过！ 做错了！？说明你对golang的方法集还有一些疑问。 一句话：golang的方法集仅仅影响接口实现和方法表达式转化，与通过实例或者指针调用方法无关。
	//var peo People = Stduent{}
	var peo Stduent = Stduent{}
	think := "bitch"
	fmt.Println(peo.Speak(think))

	t := Teacher{}
	t.ShowB()
	t.ShowA()
}