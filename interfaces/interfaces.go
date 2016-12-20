package main

import "fmt"

type animal interface {
	makeSound() string
}

type cat struct {
	name string
}

type dog struct {
	name string
}

func (c cat) makeSound() string {
	return "meeoow"
}

func (d dog) makeSound() string {
	return "woof"
}

func main() {
	c := cat{name: "guinness"}
	d := dog{name: "daisy"}

	speak(c)
	speak(d)
}

func speak(a animal) {
	fmt.Println(a.makeSound())
}
