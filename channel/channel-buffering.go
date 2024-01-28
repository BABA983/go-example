package main

import "fmt"

func error() {
	messages := make(chan string)

	messages <- "buffered"

	fmt.Println(<-messages)
}

func main() {
	messages := make(chan string, 2)

	messages <- "buffered"
	messages <- "channel"

	fmt.Println(<-messages)
	fmt.Println(<-messages)

	// error()
}
