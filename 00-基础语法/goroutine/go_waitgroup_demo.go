package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(3)
	go func() {
		defer wg.Done()
		fmt.Println("foo")
	}()
	go func() {
		defer wg.Done()
		fmt.Println("bar")
	}()
	go func() {
		defer wg.Done()
		fmt.Println("baz")
	}()
	wg.Wait()
}
