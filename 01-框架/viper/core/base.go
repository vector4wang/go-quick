package core

import "fmt"

var (
	C ViperConfig
)

func ConfigPrint() {
	fmt.Println("config print: ", C)

}
