package main

import "fmt"

func main() {
	var username string
	fmt.Println("请输入姓名!")

	_, err := fmt.Scanf("%s", &username)
	if err != nil {
		return
	}

	fmt.Printf("输入的姓名为: %s", username)

}
