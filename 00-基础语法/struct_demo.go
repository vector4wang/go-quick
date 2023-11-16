package main

import (
	"encoding/json"
	"fmt"
)

type Stu struct {
	Name string
	Age  int
	High bool
	Sex  string
}

type StuJson struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	High bool   `json:"high"`
	Sex  string `json:"sex"`
}

func main() {
	stu := Stu{
		Name: "vector",
		Age:  10,
		High: false,
		Sex:  "男",
	}

	fmt.Println(stu)
	stj, err := json.Marshal(stu)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(stj))

	stuJson := StuJson{
		Name: "vector",
		Age:  10,
		High: false,
		Sex:  "男",
	}

	sj, err := json.Marshal(stuJson)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(sj))
}
