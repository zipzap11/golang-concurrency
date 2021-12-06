package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("somehting i don't know")
	group := &sync.WaitGroup{}
	for i := 0; i < 30; i++ {
		go RunAsynchronus(group, i)
	}

	group.Wait()
	fmt.Println("finish")
}

func RunAsynchronus(group *sync.WaitGroup, i int) {
	defer group.Done()

	group.Add(1)
	fmt.Println("hellow", i)
	time.Sleep(1 * time.Second)
}
