package main

import "fmt"

type rect struct {
	width  int
	height int
}

func (r *rect) area() int {
	return r.width * r.height
}

func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func main() {
	r := rect{10, 2}

	fmt.Println(r.area())
	fmt.Println(r.perim())
	rp := &r
	fmt.Println("area: ", rp.area())
	fmt.Println("perim:", rp.perim())
}
