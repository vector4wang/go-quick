package main

import (
	"fmt"
	"math/rand"
)

func main() {
	tree1 := New(1)
	tree2 := New(2)
	fmt.Println(Same(tree1, tree2))
}

// 函数遍历树 t，将树中的所有值依次发送到通道中
func Walk(t *Tree, ch chan int) {
	if t == nil {
		return
	}
	if t.Left != nil {
		Walk(t.Left, ch)
	}
	ch <- t.Value
	if t.Right != nil {
		Walk(t.Right, ch)
	}
}

// 判断两棵树是否包含相同的值
func Same(t1, t2 *Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go Walk(t1, ch1)
	go Walk(t2, ch2)

	var count int

	for {
		if <-ch1 == <-ch2 {
			count++
			// 这里的count等于10，是因为题目要求里面随机生成的树的节点个数就是10个
			// 一般的树可以给树添加一个Len属性表示节点个数，用Len属性来判断
			if count == 10 {
				return true
			}
		} else {
			return false
		}
	}
}

// Tree 是一个包含整数值的二叉树
type Tree struct {
	Left  *Tree
	Value int
	Right *Tree
}

// New 会返回一个新的，包含值为k, 2k, ..., 10k的随机二叉树
func New(k int) *Tree {
	var t *Tree
	// Perm 返回半开区间 [0,n) 中整数的伪随机排列
	for _, v := range rand.Perm(10) {
		t = insert(t, (1+v)*k)
	}
	return t
}

// 将值插入树中，节点的左子树中存的值小于节点的值，右子树存的值大于等于节点的值
func insert(t *Tree, v int) *Tree {
	if t == nil {
		return &Tree{nil, v, nil}
	}
	if v < t.Value {
		t.Left = insert(t.Left, v)
	} else {
		t.Right = insert(t.Right, v)
	}
	return t
}

// 递归到左子树的最左边，使用递归依次拿最小的值，然后拿根节点的值，递归到右子树的最左边，使用递归依次拿最小的值
func (t *Tree) String() string {
	if t == nil {
		return "()"
	}
	s := ""
	if t.Left != nil {
		s += t.Left.String() + " "
	}
	s += fmt.Sprint(t.Value)
	if t.Right != nil {
		s += " " + t.Right.String()
	}
	return "(" + s + ")"
}
