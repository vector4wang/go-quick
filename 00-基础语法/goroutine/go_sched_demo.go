package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Printf("a %d\n", i)
			runtime.Gosched()
		}
	}()
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Printf("b %d\n", i)
			runtime.Gosched()
		}
	}()
	time.Sleep(time.Second)
}
