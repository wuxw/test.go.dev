package main

import (
	"fmt"
	"sync"
)

var x = 0

func increment(wg *sync.WaitGroup, ch chan bool) {
	ch <- true //写入值
	x = x + 1
	<-ch //读取值，这时候通道又有空位了
	wg.Done()
}
func main() {
	var w sync.WaitGroup
	//在上面的程序中，我们创建了一个容量为1的缓冲通道，这个通道被传递到直线no的增量Goroutine。这个缓冲通道用于确保只有一个Goroutine能够访问代码的关键部分，从而增加x
	ch := make(chan bool, 1)
	for i := 0; i < 1000; i++ {
		w.Add(1)
		go increment(&w, ch)
	}
	w.Wait()
	fmt.Println("final value of x", x)
}
