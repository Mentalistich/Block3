package main

import (
	"fmt"
	"sync"
)

func inChan(input []int) chan int {
	c := make(chan int)
	go func() {
		for _, v := range input {
			c <- v
		}
		close(c)
	}()
	return c
}

func fanOut(input <-chan int) <-chan int {
	c := make(chan int)
	go func() {
		for i := range input {
			c <- i * i
		}

		close(c)
	}()
	return c
}

func fanIn(channels ...<-chan int) <-chan int {
	out := make(chan int)
	wg := &sync.WaitGroup{}
	wg.Add(len(channels))
	for _, c := range channels {
		go func(channel <-chan int) {
			for r := range channel {
				out <- r
			}
			wg.Done()
		}(c)
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func main() {
	test1 := inChan([]int{1, 2, 3, 4, 5})
	test2 := inChan([]int{5, 6, 7, 8, 9})
	fanIn1 := fanOut(test2)
	fanIn2 := fanOut(test1)
	fanIn3 := fanOut(test1)

	for i := range fanIn(fanIn1, fanIn2, fanIn3) {
		fmt.Println(i)
	}
}
