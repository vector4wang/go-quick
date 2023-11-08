package main

import "fmt"

/*
*
函数多结果返回
*/
func main() {
	a, b := swap("google", "runoob")
	fmt.Println(a, b)
}

func swap(s1 string, s2 string) (string, string) {
	return s2, s1
}
