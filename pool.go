package main

import (
	"fmt"
	"sync"
)

func main() {
	group := &sync.WaitGroup{}

	// pool with default value
	pool := sync.Pool{
		New: func() interface{} {
			return "Ducky"
		},
	}

	pool.Put("data 1")
	pool.Put("data 2")
	pool.Put("data 3")

	for i := 0; i < 10; i++ {
		go func() {
			defer group.Done()

			group.Add(1)

			data := pool.Get()
			fmt.Println(data)
			pool.Put(data)
		}()
	}

	group.Wait()
}
