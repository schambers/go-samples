package main

import "fmt"

func update(s string) {
	fmt.Println(s)
}

func printMore() {
	for i := 0; i < 3; i++ {
		fmt.Println("Inside printMore")
	}
}

func main() {
	numberRoutines := 20

	fmt.Println("Entering routine loop")

	go printMore()

	for i := 1; i <= numberRoutines; i++ {
		go update(fmt.Sprintf("Running in route:%v", i))
	}

	var input string
	fmt.Scanln(&input)
}
