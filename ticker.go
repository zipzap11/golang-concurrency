package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(1 * time.Second)

	ctr := 1

	for time := range ticker.C {
		if ctr == 5 {
			ticker.Stop()
		}
		fmt.Println(time)
		ctr++
	}
}
