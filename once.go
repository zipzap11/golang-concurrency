package main

import (
	"fmt"
	"sync"
)

var counter int = 0

func main() {
	group := &sync.WaitGroup{}
	// once := &sync.Once{}

	for i := 0; i < 20; i++ {
		go func() {
			defer group.Done()
			group.Add(1)
			// once.Do(OnlyOnce)
			OnlyOnce()
		}()
	}

	group.Wait()
	fmt.Println("counter =", counter)
	fmt.Println("finish")
}

func OnlyOnce() {
	counter++
}
