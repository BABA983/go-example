package main

import (
	"fmt"
	"strconv"
	"strings"
)

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func f1(from string, c chan<- string) {
	for i := 0; i < 3; i++ {
		// fmt.Println(from, ":", i)
		c <- strings.Join([]string{from, strconv.Itoa(i)}, ":")
	}
	close(c)
}

func main() {
	// messages := make(chan string)
	messages := make(chan string, 2)

	go func() { messages <- "ping" }()

	f("hahahah")

	go f1("async", messages)

	// go func() { messages <- strings.Join([]string{"async", "1"}, "::") }()

	for msg := range messages {
		fmt.Println(msg)
	}
	// fmt.Println(<-messages)
	// fmt.Println(<-messages)
	// fmt.Println(<-messages)
	// fmt.Println(<-messages)
}
