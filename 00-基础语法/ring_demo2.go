package main

import (
	"container/ring"
	"fmt"
)

const n = 6
const m = 5

func main() {
	r := ring.New(n)
	// 给环填充值
	for i := 1; i <= n; i++ {
		r.Value = i
		r = r.Next()
	}
	cnt := 1
	for r.Len() > 1 {
		r = r.Move(m - 2)
		fmt.Printf("第%d次淘汰的编号为%d\n", cnt, r.Next().Value)
		r.Unlink(1)
		r = r.Next()
		cnt++
	}
	fmt.Println("最终结果为", r.Value)
}
