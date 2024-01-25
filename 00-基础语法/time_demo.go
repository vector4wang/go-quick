package main

import (
	"fmt"
	"time"
)

func main() {

	t1, err := time.Parse("2006-01-02 15:04:05", "2022-12-03 13:00:00")
	if err != nil {
		fmt.Println("err: ", err)
		return
	}
	fmt.Println(t1) // 2022-12-03 13:00:00 +0000 UTC

	t2, err := time.Parse("2006-01-02", "2022-12-03")
	if err != nil {
		fmt.Println("err: ", err)
		return
	}
	fmt.Println(t2) // 2022-12-03 00:00:00 +0000 UTC

	t3, err := time.Parse("15:04:05", "13:00:00")
	if err != nil {
		fmt.Println("err: ", err)
		return
	}
	fmt.Println(t3) // 0000-01-01 13:00:00 +0000 UTC

	t1 = time.Now()
	fmt.Println(t1)
	time.Sleep(time.Duration(2) * time.Second)
	t2 = time.Now()
	fmt.Println(t2.Sub(t1))

}
