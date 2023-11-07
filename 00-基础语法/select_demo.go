package main

import (
	"fmt"
	"time"
)

/*
*
select 是 Go 中的一个控制结构，类似于 switch 语句。
select 语句只能用于通道操作，每个 case 必须是一个通道操作，要么是发送要么是接收。
select 语句会监听所有指定的通道上的操作，一旦其中一个通道准备好就会执行相应的代码块。
如果多个通道都准备好，那么 select 语句会随机选择一个通道执行。如果所有通道都没有准备好，那么执行 default 块中的代码。

每个 case 都必须是一个通道
所有 channel 表达式都会被求值
所有被发送的表达式都会被求值
如果任意某个通道可以进行，它就执行，其他被忽略。
如果有多个 case 都可以运行，select 会随机公平地选出一个执行，其他不会执行。
否则：

	如果有 default 子句，则执行该语句。
	如果没有 default 子句，select 将阻塞，直到某个通道可以运行；Go 不会重新对 channel 或值进行求值。
*/
func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received ", msg1)
		case msg2 := <-c2:
			fmt.Println("received ", msg2)

		}
	}
}
