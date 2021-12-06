package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var counter int32 = 1
	group := sync.WaitGroup{}

	for i := 0; i < 20; i++ {
		group.Add(1)

		func() {
			atomic.AddInt32(&counter, 1)
		}()

		group.Done()
	}

	group.Wait()

	fmt.Println("result =", counter)
}
