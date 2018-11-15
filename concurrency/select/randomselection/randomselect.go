package main

import (
	"fmt"
	"time"
	//"time"
)

func server1(ch chan string) {
	ch <- "from server1"
}
func server2(ch chan string) {
	ch <- "from server2"

}

func main() {
	output1 := make(chan string)
	output2 := make(chan string)
	go server1(output1)
	go server2(output2)
	/*
			在上面的程序中，server1和server2 go例程被称为行no。分别是18和19。然后主程序在第一行中休眠1秒。20.当控制到达select语句的时候。server1将从server1写到output1通道，server2将从server2写到output2通道，因此select语句的情况都准备好执行了。如果您多次运行这个程序，那么输出将从server1或server2之间变化，这取决于随机选择的情况。
		请在您的本地系统中运行这个程序以获得这种随机性。如果这个程序在操场上运行，它将打印相同的输出，因为操场是确定的。
	*/
	time.Sleep(1 * time.Second) //如果增加这个那么server1和server2都好了，程序会随机执行case 反之如果没有就只会匹配第一个case
	select {
	case s1 := <-output1:
		fmt.Println(s1)
	case s2 := <-output2:
		fmt.Println(s2)
	}
}
