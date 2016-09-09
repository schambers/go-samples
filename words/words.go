package main

import (
	"fmt"
	"strings"
)

func main() {
	wordsToReverse := "This is a list of words in a sentence"
	wordsArray := strings.Split(wordsToReverse, " ")
	fmt.Printf("Length: %v\n\n", len(wordsArray))

	for i := 0; i < len(wordsArray); i++ {
		fmt.Printf("Word entry:%v\n", wordsArray[i])
	}
}
