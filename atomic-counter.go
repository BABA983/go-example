package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	var ops atomic.Uint64

	var wg sync.WaitGroup

	for i := 0; i < 50; i++ {
		wg.Add(1)

		go func() {
			for c := 0; c < 1000; c++ {
				ops.Add(1)
			}

			wg.Done()
		}()
	}

	j := 0

	for i := 0; i < 50; i++ {
		go func() {
			for c := 0; c < 1000; c++ {
				j++
			}
		}()
	}

	wg.Wait()
	time.Sleep(2 * time.Second)

	fmt.Println("j", j)
	fmt.Println("ops:", ops.Load())
}
