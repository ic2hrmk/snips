package main

import (
	"fmt"
	"math/rand"
	"reflect"
	"sync"
	"time"
)

func main() {
	myChan1 := asChan(1, 2, 3, 4, 5)
	myChan2 := asChan(6, 7, 8, 9, 10)
	myChan3 := asChan(11, 12, 13, 14, 15)

	//for v := range merge(myChan1, myChan2, myChan3) {
	//	fmt.Println(v)
	//}

	for v := range mergeReflect(myChan1, myChan2, myChan3) {
		fmt.Println(v)
	}
}

func merge(chans ...<-chan int) <-chan int {
	merged := make(chan int)

	go func() {
		var wg sync.WaitGroup
		wg.Add(len(chans))

		for _, c := range chans {
			go func(c <-chan int) {
				for v := range c {
					merged <- v
				}
				wg.Done()
			}(c)
		}
		wg.Wait()
		close(merged)
	}()

	return merged
}

func mergeReflect(chans ...<-chan int) <-chan int {
	merged := make(chan int)

	go func() {
		defer close(merged)
		cases := []reflect.SelectCase{}

		for _, c := range chans {
			cases = append(cases, reflect.SelectCase{
				Dir:  reflect.SelectRecv,
				Chan: reflect.ValueOf(c),
			})
		}

		for len(cases) > 0 {
			i, v, ok := reflect.Select(cases)
			if !ok {
				cases = append(cases[:i], cases[i+1:]...)
				continue
			}

			merged <- v.Interface().(int)
		}
	}()

	return merged
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
