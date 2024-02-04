package main

import (
	"fmt"
	"time"
)

func main() {
	var count int
	//var mu sync.Mutex
	for i := 0; i < 100; i++ {
		go func() {
			//mu.Lock()
			count++
			//mu.Unlock()
		}()
	}
	time.Sleep(time.Second)
	fmt.Println(count)
}
