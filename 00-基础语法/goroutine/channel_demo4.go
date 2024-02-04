package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 1000; i++ {
		//go func(n int) {
		//	fmt.Println(n)
		//}(i)

		go func() {
			fmt.Println(i)
		}()

	}
	time.Sleep(time.Second)
}
