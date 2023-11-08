package main

import (
	"fmt"
)

type usber interface {
	start()
	stop()
}
type Computer struct {
	name  string
	model string
}

func (c Computer) start() {
	fmt.Println("打开电脑", c.name)
}
func (c Computer) stop() {
	fmt.Println("关闭电脑", c.name)
}

func (c Computer) FrontPlay() {
	fmt.Println("前置音乐", c.name)
}

func (c Computer) BackPlay() {
	fmt.Println("后置音乐", c.name)
}

type Phone struct {
	name  string
	model string
}

func (p Phone) start() {
	fmt.Println("启动手机")
}

func (p Phone) stop() {
	fmt.Println("关闭手机")
}
func working(u usber) {
	u.start()
	u.stop()
}

func PlayAuto(a Audio) {
	a.FrontPlay()
	a.BackPlay()
}
func main() {
	cm := Computer{name: "mac", model: "A3"}
	working(cm)
	PlayAuto(cm)
	//Audio(cm).BackPlay()
}
