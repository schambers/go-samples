package main

import (
	"fmt"
)

type cat struct {
	name string
}

func (c cat) makeSound() string {
	return "meoowww"
}

func main() {
	c := cat{name: "guinness"}
	fmt.Printf("The cat says %s\n", c.makeSound())
}
