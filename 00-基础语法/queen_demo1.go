package main

import (
	"container/list"
	"fmt"
)

// 先进先出
type student1 struct {
	name string
	age  int
}

func Push1(l *list.List, v interface{}) *list.Element {
	return l.PushBack(v)
}

func Pop1(l *list.List) interface{} {
	elem := l.Front()
	return l.Remove(elem)
}

func main() {
	// 初始化栈
	stack := list.New()
	// 栈的push操作
	Push1(stack, student1{name: "dexuan", age: 3})
	Push1(stack, student1{name: "dexuan2", age: 4})
	Push1(stack, student1{name: "dexuan2", age: 5})
	// 从头开始遍历
	for head := stack.Front(); head != nil; head = head.Next() {
		fmt.Println(head.Value)
	}
	//栈的pop操作
	s := Pop1(stack)
	fmt.Println("pop:", s)
	// 从头开始遍历
	for head := stack.Front(); head != nil; head = head.Next() {
		fmt.Println(head.Value)
	}
}
