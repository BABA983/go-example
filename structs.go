package main

import "fmt"

type person struct {
	name string
	age  int
}

func getPerson(name string) *person {
	p := person{name, 24}
	return &p
}

func main() {
	fmt.Println(person{"who", 18})
	fmt.Println(person{name: "whoWithNoAge"})
	fmt.Println(&person{name: "whoWithNoAge"})
	p := getPerson("David")
	fmt.Println(p.name)
	p1 := person{"Sam", 44}
	pp := &p1
	fmt.Println(pp.age)
	fmt.Println(getPerson("haha"))
	dog := struct {
		name   string
		isGood bool
	}{
		"Rex",
		true,
	}
	fmt.Println(dog)
}
