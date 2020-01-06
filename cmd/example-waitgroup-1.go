package main

import (
    "fmt"
    "sync"
    "time"
)

func main() {
    var wg sync.WaitGroup
    fmt.Printf("init:             %+v\n", wg)

    for i := 1; i < 10; i++ {
        // 计数加 1
		wg.Add(1)  // Add 用来添加 goroutine 的个数
		// wg.Add() 方法一定要在 goroutine 开始前执行
        go func(i int) {
            fmt.Printf("goroutine%d start: %+v\n", i, wg)
            time.Sleep(11 * time.Second)
            // 计数减 1
            wg.Done()  //Done 执行一次数量减 1
            fmt.Printf("goroutine%d end:   %+v\n", i, wg)
        }(i)
        time.Sleep(time.Second)
    }

    wg.Wait() //Wait 用来等待结束
    fmt.Printf("over:             %+v\n", wg)
}