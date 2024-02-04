package main

import (
	"fmt"
	"time"
)

func main() {
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)
	var x, y time.Time
	for {
		select {
		case x, _ = <-tick:
			fmt.Println(x, "tick.")
		case y, _ = <-boom:
			fmt.Println(y, "BOOM!")
			return
		default:
			fmt.Println("    .")
			time.Sleep(50 * time.Millisecond)
		}
	}
}
