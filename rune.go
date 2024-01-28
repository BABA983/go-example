package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	const s = "hello"
	const ss = "你好"

	fmt.Println("len", len(s))
	fmt.Println("Rune count:", utf8.RuneCountInString(ss))

	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", ss[i])
	}

	fmt.Println()

	for idx, runeValue := range ss {
		fmt.Printf("%#U starts at %d\n", runeValue, idx)
	}
}
