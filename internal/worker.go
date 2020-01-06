package jobqueue

import (
    "fmt"
	"sync"
)




// make a channel with a capacity of 100.
//jobChan := make(chan Job, 100)

// start the worker
//go worker(jobChan)

// enqueue a job
//jobChan <- job