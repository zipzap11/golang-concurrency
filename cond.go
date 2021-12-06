package main

import (
	"fmt"
	"sync"
	"time"
)

var cond = sync.NewCond(&locker)
var locker = sync.Mutex{}
var group = sync.WaitGroup{}

func main() {
	for i := 1; i <= 10; i++ {
		go WaitCondition(i)
	}

	// go func() {
	// 	for i := 1; i <= 10; i++ {
	// 		time.Sleep(1 * time.Second)
	// 		cond.Signal()
	// 	}
	// }()

	go func() {
		time.Sleep(1 * time.Second)
		cond.Broadcast()
	}()

	group.Wait()
}

func WaitCondition(value int) {
	defer group.Done()
	group.Add(1)

	cond.L.Lock()
	cond.Wait()
	fmt.Println("Done", value)
	cond.L.Unlock()
}
