package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	group := sync.WaitGroup{}
	group.Add(1)
	Timer()
	After()
	afterFunc(&group)
	group.Wait()
}

func Timer() {
	timer := time.NewTimer(3 * time.Second)
	fmt.Println(time.Now())

	time := <-timer.C
	fmt.Println(time)
}

func After() {
	channel := time.After(3 * time.Second)
	fmt.Println(time.Now())

	time := <-channel
	fmt.Println(time)
}

func afterFunc(group *sync.WaitGroup) {
	time.AfterFunc(3*time.Second, func() {
		fmt.Println("run after 3")
		fmt.Println(time.Now())
		group.Done()
	})
	fmt.Println("pass")
	fmt.Println(time.Now())
}
