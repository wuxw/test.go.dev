package main

import (
	"fmt"
	"sync"
)

//全局变量
var x = 0

func increment(wg *sync.WaitGroup, m *sync.Mutex) {
	//func increment(wg *sync.WaitGroup) {
	m.Lock()
	x = x + 1
	m.Unlock()
	wg.Done()
}

/*

我们已经用互斥体和通道解决了竞态条件问题。那么我们如何决定什么时候使用。答案就在你想要解决的问题上。如果您想要解决的问题更适合互斥锁，那么就继续使用互斥锁。如果需要，不要犹豫使用互斥锁。如果问题似乎更适合于渠道，那么就使用它：）。
大多数Go新手都试图用一个通道来解决所有并发问题，因为这是该语言的一个很酷的特性。这是错误的。该语言为我们提供了使用互斥或通道的选项，并且在选择中也没有错误。
当Goroutines需要相互通信时，当只有一个Goroutine应该访问关键的代码段时，通常使用通道。
在我们上面解决的问题的情况下，我更倾向于使用互斥锁，因为这个问题不需要在goroutines之间进行任何通信。因此，互斥锁是很自然的。
*/
//Mutex更适合没有通信的时候，来解决竞态条件，channel更适合有通信的时候来解决竞态条件
func main() {
	var w sync.WaitGroup
	var m sync.Mutex
	for i := 0; i < 1000; i++ {
		w.Add(1)
		//go increment(&w)
		go increment(&w, &m)
	}
	w.Wait()
	fmt.Println("final value of x", x)
}
