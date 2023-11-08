package main

import "fmt"

type Circle struct {
	radius float64
}

func main() {
	c1 := Circle{
		radius: 10.00,
	}

	fmt.Println("圆的面积", c1.getArea())
}

// 该 method 属于 Circle 类型对象中的方法
func (c Circle) getArea() float64 {
	return 3.14 * c.radius * c.radius
}
