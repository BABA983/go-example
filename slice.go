package main

import "fmt"

func main() {
	s := make([]string, 3)
	fmt.Println("example:", s)

	s[0] = "a"
	s[1] = "c"
	fmt.Println("len:", len(s))

	s = append(s, "b")
	fmt.Println("example:", s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", len(c), c)

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)
}
