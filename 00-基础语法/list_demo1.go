package main

import "fmt"

func main() {
	// 切片
	array := []int{1, 2, 3}
	fmt.Println(array)

	array = append(array, 5, 6, 7, 8, 9)
	fmt.Println(array)

	// 数组
	array2 := [4]int{1, 2, 3}
	fmt.Println(array2)

}
