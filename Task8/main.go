package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(i int, c <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for n := range c {
		//j := <-c
		fmt.Printf("Worker %d выполняет задачу %d\n", i, n)
		time.Sleep(1 * time.Second)
		fmt.Printf("Worker %d выполнил задачу %d\n", i, n)
	}
}

func main() {
	ch := make(chan int)
	var wg sync.WaitGroup

	for i := 1; i < 6; i++ {
		wg.Add(1)
		go worker(i, ch, &wg)
	}

	for j := 1; j < 6; j++ {
		ch <- j
	}
	close(ch)
	wg.Wait()
	fmt.Println("Дело сделано")
}
