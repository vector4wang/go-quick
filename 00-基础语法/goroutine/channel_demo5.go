package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("Worker", id, "processing job", j)
		time.Sleep(time.Second)
		results <- j * 2
	}
}

func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	// 启动3个Goroutine进行工作
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// 发送5个任务到通道中
	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)

	// 输出处理结果
	for a := 1; a <= 5; a++ {
		<-results
	}
}
