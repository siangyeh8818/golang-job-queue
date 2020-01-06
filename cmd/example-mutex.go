package main

import (
    "fmt"
	"sync"
	"time"
)

func main() {
    var mutex sync.Mutex
    fmt.Printf("%+v\n", mutex)
	mutex.Lock()
	
	for i:=0;i<10;i++ {
		if i ==8 {
			mutex.Unlock()
		}
		fmt.Printf("i = %d\n", i)
		fmt.Printf("%+v\n", mutex)
		time.Sleep(1 * time.Second)
	}
    

    //mutex.Unlock()
    //fmt.Printf("%+v\n", mutex)
}