package main

import (
    "fmt"
	"sync"
	"time"
	//. "github.com/siangyeh8818/golang-job-queue/internal"
)

var jobChan chan int
var wg sync.WaitGroup

func main() {
	jobChan = make(chan int, 100)
	
    for i := 1; i <= 10; i++{
        jobChan <- i
    }

    wg.Add(1)
    close(jobChan)

	go worker(jobChan , &wg)
	res := WaitTimeout(&wg, 5*time.Second)
    if res {
        fmt.Println("执行完成退出")
    } else {
        fmt.Println("执行超时退出")
    }
    //wg.Wait()
}

func worker(jobChan <- chan int , wg *sync.WaitGroup)  {
    defer wg.Done()
    for job := range jobChan{
		//process(job)
		fmt.Printf("执行任务 %d \n", job)
		time.Sleep(1 * time.Second)
    }
}
	
	//超时机制
func WaitTimeout(wg *sync.WaitGroup, timeout time.Duration) bool {
    ch := make(chan struct{})
    go func() {
        wg.Wait()
        close(ch)
    }()
    select {
    case <-ch:
        return true
    case <-time.After(timeout):
        return false
    }
}