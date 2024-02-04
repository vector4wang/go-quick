package main

import (
	"fmt"
	"sync"
	"time"
)

var counter int
var mutex sync.Mutex

func increment() {
	//mutex.Lock()
	counter++
	//mutex.Unlock()
}

func main() {
	for i := 0; i < 5; i++ {
		go increment()
	}

	time.Sleep(time.Second)
	fmt.Println("Counter:", counter)
}
