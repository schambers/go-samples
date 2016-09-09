package main

import "fmt"

func main() {

	messages := make(chan string, 4)

	messages <- "buffered"
	messages <- "channel"
	messages <- "buffer 3"
	messages <- "buffer 4"

	fmt.Println(<-messages)
	fmt.Println(<-messages)
	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
