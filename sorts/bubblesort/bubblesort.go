package main

import "fmt"

func main() {
	fmt.Println("Bubble sort")

	sort := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}

	bubbleSort(sort)
}

func bubbleSort(sort []int) {
	fmt.Printf("Start: %v\n", sort)

	length := len(sort)

	// continue looping until we stop swapping items
	moved := true
	for ok := true; ok; ok = moved {
		moved = false
		for i := 1; i <= length-1; i++ {
			// Check if this pair is out of order
			if sort[i-1] > sort[i] {
				temp := sort[i-1]
				sort[i-1] = sort[i]
				sort[i] = temp

				moved = true
			}
			fmt.Printf("%v\n", sort)
		}
	}

	fmt.Printf("End: %v\n", sort)
}
