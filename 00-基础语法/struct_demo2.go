package main

import "fmt"

// 定义点结构
type Point struct {
	X int
	Y int
}

// 非指针接收器的加方法
func (p Point) Add(other Point) Point {
	// 成员值与参数相加后返回新的结构
	return Point{p.X + other.X, p.Y + other.Y}
}
func main() {
	// 初始化点
	p1 := Point{1, 1}
	p2 := Point{2, 2}
	// 与另外一个点相加
	result := p1.Add(p2)
	// 输出结果
	fmt.Printf("%v", p1)
	fmt.Printf("%v", p2)
	fmt.Println(result)
}
