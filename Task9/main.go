package main

import (
	"fmt"
)

func increment(ch <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for i := range ch {
			out <- i + 1
		}
		close(out)
	}()
	return out
}

func multiple(ch <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for i := range ch {
			out <- i * 2
		}
		close(out)
	}()

	return out
}

func toString(ch <-chan int) <-chan string {
	out := make(chan string)
	go func() {
		for i := range ch {
			incremented := i / 2
			initial := incremented - 1
			out <- fmt.Sprintf("Initial value: %d -> Incremented: %d -> Multiplied: %d", initial, incremented, i)
		}
		close(out)
	}()

	return out

}

func main() {
	in := make(chan int)
	turn1 := increment(in)
	turn2 := multiple(turn1)
	turn3 := toString(turn2)
	go func() {
		for i := 1; i < 5; i++ {
			in <- i
		}
		close(in)
	}()
	for printChan := range turn3 {
		fmt.Println(printChan)
	}
}
