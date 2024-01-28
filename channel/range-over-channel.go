package main

import "fmt"

func main() {
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)
	// queue <- "3"

	// before you iterate the channel you will need to close it
	// otherwise you will meet a dead lock
	for item := range queue {
		fmt.Println(item)
	}
}
