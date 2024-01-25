package main

import "fmt"

func main() {
	scene := make(map[string]int)
	scene["cat"] = 66
	scene["dog"] = 4
	scene["pig"] = 960
	for k, v := range scene {
		fmt.Println(k, v)
	}

}
