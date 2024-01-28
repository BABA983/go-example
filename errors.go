package main

import (
	"errors"
	"fmt"
)

func f1(num int) (int, error) {
	if num%2 == 0 {
		return -1, errors.New("can't % 2 clearly")
	}
	return num, nil
}

type myError struct {
	code int
	msg  string
}

func (e *myError) Error() string {
	return fmt.Sprintf("%d - %s", e.code, e.msg)
}

func f2(num int) (int, error) {
	if num%2 == 0 {
		return -1, &myError{-1, "error"}
	}
	return num, nil
}

func main() {
	arr := []int{1, 2, 3, 4}
	for i, k := range arr {
		fmt.Println("for loop", i, k)
		if v, err := f1(i); err != nil {
			fmt.Println("failed:", err)
		} else {
			fmt.Println(v)
		}
	}
	for _, i := range []int{7, 42} {
		if r, e := f2(i); e != nil {
			fmt.Println("f2 failed:", e)
		} else {
			fmt.Println("f2 worked:", r)
		}
	}

	// get the error as a instance
	_, e := f2(42)
	if ee, ok := e.(*myError); ok {
		fmt.Println(ee.code)
		fmt.Println(ee.msg)
	}
}
