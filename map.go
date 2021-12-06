package main

import (
	"fmt"
	"sync"
)

func main() {
	data := sync.Map{}
	group := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		go AddToMap(&data, &group, i)
	}

	group.Wait()

	data.Range(func(key, value interface{}) bool {
		fmt.Println("key =", key, "value =", value)
		return true
	})
	fmt.Println("finish")
}

func AddToMap(container *sync.Map, group *sync.WaitGroup, value int) {
	defer group.Done()
	group.Add(1)

	fmt.Println("value =", value)
	container.Store(value, value)
}
