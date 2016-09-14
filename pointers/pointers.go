package main

import "fmt"

func modifyInt(i int) {
	i = 42
	fmt.Printf("modifyInt value:%v\n", i)
}

func modifyPointer(i *int) {
	*i = 56
	fmt.Printf("modifyPointer value:%v\n", *i)
}

func main() {
	num := 43
	fmt.Printf("before modifyInt:%v\n", num)
	modifyInt(num)
	fmt.Printf("after modifyInt:%v\n", num)
	modifyPointer(&num)
	fmt.Printf("after modifyPointer:%v\n", num)
}
