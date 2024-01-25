package main

import (
	"fmt"
)

type Student struct {
	Name string
	Age  int
	High bool
	Sex  string
}

func NewStudent(name string) *Student {
	return &Student{
		Name: name,
		Age:  10,
		High: false,
		Sex:  "男",
	}
}

func (stu *Student) SetName(name string) {
	stu.Name = name
}
func (stu Student) GetName() string {
	return stu.Name
}

func main() {
	stu := NewStudent("张三")
	fmt.Println("%-v", stu)
	fmt.Println(stu.GetName())
}
