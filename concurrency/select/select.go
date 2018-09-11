package main

import (
	"fmt"
	"time"
)

func server1(ch chan string) {
	time.Sleep(6 * time.Second)
	//赋值
	ch <- "from server1"
}
func server2(ch chan string) {
	time.Sleep(3 * time.Second)
	ch <- "from server2"

}

func process(ch chan string) {
	time.Sleep(10500 * time.Millisecond)
	//time.Sleep(1000 * time.Millisecond)
	ch <- "process successful"
}

func main() {
	output1 := make(chan string)
	output2 := make(chan string)
	go server1(output1)
	go server2(output2)
	/**
	让我们假设我们有一个任务关键的应用程序，我们需要尽快将输出返回给用户。这个应用程序的数据库被复制并存储在世界各地的不同服务器中。假设功能server1和server2实际上是与两台这样的服务器通信的。每个服务器的响应时间依赖于每个服务器上的负载和网络延迟。我们将请求发送到两个服务器，然后使用select语句等待相应的响应通道。首先响应的服务器由select选择，而另一个响应被忽略。通过这种方式，我们可以将相同的请求发送到多个服务器，并返回对用户的最快响应：）。
	*/
	//select 执行了一个后，就会跳出，会跳出这个代码 继续执行
	select {
	//读取

	case s1 := <-output1:
		fmt.Println(s1)
	case s2 := <-output2:
		fmt.Println(s2)

	}

	ch := make(chan string)
	go process(ch)
	for {
		time.Sleep(1000 * time.Millisecond) //主要是怕跑的太快了，用时间间隔隔开，让开发者可以观察
		select {
		case v := <-ch:
			fmt.Println("received value: ", v)
			return
		default:
			fmt.Println("no value received")
		}
	}
}
