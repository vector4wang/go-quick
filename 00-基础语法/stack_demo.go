package main

import (
	"container/list"
	"fmt"
)

// 后进先出
type std struct {
	name string
	age  int
}

func Push(l *list.List, v interface{}) *list.Element {
	return l.PushBack(v)
}

func Pop(l *list.List) interface{} {
	elem := l.Back()
	return l.Remove(elem)
}

func main() {
	// 初始化栈
	stack := list.New()
	// 栈的push操作
	Push(stack, std{name: "dexuan", age: 3})
	Push(stack, std{name: "dexuan2", age: 4})
	Push(stack, std{name: "dexuan2", age: 5})
	// 从头开始遍历
	for head := stack.Front(); head != nil; head = head.Next() {
		fmt.Println(head.Value)
	}
	//栈的pop操作
	s := Pop(stack)
	fmt.Println("pop:", s)
}
