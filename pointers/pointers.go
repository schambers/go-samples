package main

import "fmt"

func updateValue(val *int) {
	*val = *val + 100
}

func main() {
	val := 1000
	updateValue(&val)
	fmt.Println("val:", val)
}
