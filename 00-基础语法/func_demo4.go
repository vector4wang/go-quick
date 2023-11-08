package main

import (
	"fmt"
	"math"
)

/**
Go 语言可以很灵活的创建函数，并作为另外一个函数的实参
*/

func main() {
	/* 声明函数变量 */
	getSquareRoot := func(x float64) float64 {
		return math.Sqrt(x)
	}

	/* 使用函数 */
	fmt.Println(getSquareRoot(9))

}
