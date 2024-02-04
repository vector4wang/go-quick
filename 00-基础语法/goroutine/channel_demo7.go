package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}
	fmt.Println("main end")
}

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
		time.Sleep(time.Second * 5)
	}
	// 必须在遍历结束之后关闭通道
	// 否则 for i := range c 会一直等待通道关闭
	// 只有发送方可以关闭一个通道，接收方不可以。在一个已经关闭的通道上进行发送会导致一个错误（panic）。
	// 通道不像文件，不需要总是关闭它们。关闭只有必须告诉接收方不会再来更多值时，才是必须的，比如终止一个
	close(c)
}
