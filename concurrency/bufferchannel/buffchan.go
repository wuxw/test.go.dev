package main

import (
	"fmt"
	"sync"
	"time"
)

func process(i int, wg *sync.WaitGroup) {
	fmt.Println("started Goroutine ", i)
	time.Sleep(2 * time.Second)
	fmt.Printf("Goroutine %d ended\n", i)
	wg.Done()
}

func write(ch chan int) {
	for i := 0; i < 5; i++ {
		ch <- i
		fmt.Println("successfully wrote", i, "to ch")
	}
	close(ch)
}

func main() {
	/* ch1 := make(chan string, 2)
	ch1 <- "naveen"
	ch1 <- "paul"
	//ch1 <- "steve" //Deadlock
	fmt.Println(<-ch1)
	fmt.Println(<-ch1)
	//ch1 <- "steve" //这里不会Deadlock
	//Deadlock
	//fmt.Println(<-ch1) */

	no := 3
    var wg sync.WaitGroup
    for i := 0; i < no; i++ {
        wg.Add(1)
        go process(i, &wg)
    }
    wg.Wait()
    fmt.Println("All go routines finished executing")

	ch1 := make(chan string, 2)
	ch1 <- "naveen"
	ch1 <- "paul"
	fmt.Println("capacity is", cap(ch1))
	fmt.Println("length is", len(ch1))
	fmt.Println("read value", <-ch1)
	fmt.Println("new length is", len(ch1))

	/* ch := make(chan int, 2)
	go write(ch)
	time.Sleep(2 * time.Second)
	for v := range ch {
		fmt.Println("read value", v, "from ch")
		time.Sleep(2 * time.Second)

	} */

}
