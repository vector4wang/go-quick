package main

import (
	"container/ring"
	"fmt"
)

func main() {
	// 初始化3个元素的环，返回头节点
	r := ring.New(3)

	// 给环填充值
	for i := 1; i <= 3; i++ {
		r.Value = i
		r = r.Next()
	}
	sum := 0
	// 对环的每个元素进行处理
	r.Do(func(i interface{}) {
		sum = i.(int) + sum
	})
	fmt.Println(sum)

}
