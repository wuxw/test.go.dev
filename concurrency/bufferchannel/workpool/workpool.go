package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Job struct {
	id       int
	randomno int
}

type Result struct {
	job         Job
	sumofdigits int
}

func digits(number int) int {
	sum := 0
	no := number
	for no != 0 {
		digit := no % 10
		sum += digit
		no /= 10
	}

	time.Sleep(2 * time.Second)
	return sum

}

var jobs = make(chan Job, 10)
var results = make(chan Result, 10)

func worker(wg *sync.WaitGroup) {
	for job := range jobs {
		//fmt.Println(job) 	//jobs的赋值来自于job通道
		result := Result{job, digits(job.randomno)}
		fmt.Println(" worker for")
		//result 赋值
		results <- result
	}
	fmt.Println(" worker done")
	wg.Done()
}

func createWorkerPool(noOfWorkers int) {
	var wg sync.WaitGroup
	for i := 0; i < noOfWorkers; i++ {
		wg.Add(1)
		fmt.Println(" createWorkerPool for")
		go worker(&wg)
	}
	wg.Wait()
	fmt.Println(" createWorkerPool close")
	close(results)
}

func allocate(noOfJobs int) {
	for i := 0; i < noOfJobs; i++ {
		randomno := rand.Intn(999)
		job := Job{i, randomno}
		//jobs的赋值来自于job通道
		jobs <- job
		fmt.Println(" allocate for")
	}
	fmt.Println(" allocate close")
	close(jobs)
}

func result(done chan bool) {
	for result := range results {
		fmt.Printf("Job id %d, input random no %d , sum of digits %d\n", result.job.id, result.job.randomno, result.sumofdigits)
	}
	done <- true
}

func main() {
	startTime := time.Now()
	noOfJobs := 12 //作业投递
	go allocate(noOfJobs)
	done := make(chan bool)
	go result(done)
	noOfWorkers := 8 //工作处理进程 越大执行耗时越少
	createWorkerPool(noOfWorkers)
	<-done
	endTime := time.Now()
	diff := endTime.Sub(startTime)
	fmt.Println("total time taken ", diff.Seconds(), "seconds")
}
