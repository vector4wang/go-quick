package main

import (
	"container/list"
	"fmt"
)

type student struct {
	name string
	age  int
}

func main() {
	// 初始化双向链表
	l := list.New()
	// 链表头插入
	l.PushFront(student{name: "dexuan", age: 3})
	// 链表尾插入
	l.PushBack(student{name: "dexuan2", age: 4})
	l.PushFront(student{name: "dexuan3", age: 5})
	// 从头开始遍历
	for head := l.Front(); head != nil; head = head.Next() {
		fmt.Println(head.Value)
	}

}
