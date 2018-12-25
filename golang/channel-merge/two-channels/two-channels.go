package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	myChan1 := asChan(1, 2, 3, 4, 5)
	myChan2 := asChan(6, 7, 8, 9, 10)

	c := mergeChans(myChan1, myChan2)

	for v := range c {
		fmt.Println(v)
	}
}

func asChan(in ...int) <-chan int {
	c := make(chan int)

	go func() {
		for i := range in {
			c <- in[i]
			time.Sleep(time.Duration(rand.Int63n(1000)) * time.Millisecond)
		}
		close(c)
	}()

	return c
}

func mergeChans(a, b <-chan int) <-chan int {
	merged := make(chan int)

	go func() {
		defer close(merged)

		for a != nil || b != nil {

			select {
			case v, ok := <-a:
				if !ok {
					a = nil
					continue
				}

				merged <- v

			case v, ok := <-b:
				if !ok {
					b = nil
					continue
				}

				merged <- v
			}
		}
	}()

	return merged
}
