package main

import (
    "fmt"
    "sync"
    "time"
)

type SafeCounter struct {
    v   map[string]int
    mux sync.Mutex
}

func (c *SafeCounter) Inc(key string) {
    c.mux.Lock()
    time.Sleep(3 * time.Second)
    c.v[key]++
    c.mux.Unlock()
}

func main() {
    c := SafeCounter{v: make(map[string]int)}
    go c.Inc("somekey")

    time.Sleep(1 * time.Second)
    c.v["somekey"]++
    fmt.Println(c.v["somekey"])

    time.Sleep(3 * time.Second)
    fmt.Println(c.v["somekey"])
}