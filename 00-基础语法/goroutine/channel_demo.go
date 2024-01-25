package main

import "fmt"

// 发送和接收必须成对出现
func main() {
	ch := make(chan int)
	go func() {
		ch <- 0
		ch <- 1
		ch <- 2
		ch <- 3
	}()

	data, ok := <-ch //没有值出来会阻塞

	fmt.Println(data)
	fmt.Println(ok)

	// 有缓冲
	chb := make(chan int, 1) // 创建一个容量为1的有缓冲区通道
	chb <- 10
	fmt.Println("发送成功")

	// 关闭
	c := make(chan int)
	go func() {
		for i := 0; i < 5; i++ {
			c <- i
		}
		close(c)
	}()
	for {
		if data, ok := <-c; ok {
			fmt.Println(data)
		} else {
			break
		}
	}
	fmt.Println("main结束")

}
