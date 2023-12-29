package main

import (
	"fmt"
)

type carListType map[string]string

var carList = make(carListType)

func main() {
	age := 10
	fmt.Printf("addr is:%p\n", &age) //addr is:0xc000018088
	sayAge(age)

	setAge(&age)
	fmt.Printf("after setAge, age is:%d\n", age) // after setAge, age is:30

	carList["honda"] = "civic"
	carList["bmw"] = "320li"

	fmt.Printf("carList is:%v\n", carList)              // carList is:map[bmw:320li honda:civic]
	fmt.Printf("carList value is:%p\n", carList)        // carList value is:0xc000098000
	fmt.Printf("carList addr is:%p\n", &carList)        // carList addr is:0x1173648
	setCar(carList)                                     // setCar carList addr is:0xc00008e000
	fmt.Printf("after setCar carList is:%v\n", carList) // after setCar carList is:map[bmw:520li honda:civic]
}

func sayAge(age int) {
	fmt.Printf("addr is:%p\n", &age)  //addr is:0xc000018098
	fmt.Printf("my age is:%d\n", age) // after setAge, age is:30
}

func setAge(age *int) {
	*age = 30
	fmt.Printf("age point value is:%p\n", age) //age point value is:0xc000080008
	fmt.Printf("age point addr is:%p\n", &age) //age point addr is:0xc00008a020
}

func setCar(carList carListType) {
	fmt.Printf("setCar carList value is:%p\n", carList) // setCar carList value is:0xc000094000
	fmt.Printf("setCar carList addr is:%p\n", &carList) // setCar carList addr is:0xc00008e020
	carList["bmw"] = "520li"
}
