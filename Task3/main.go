package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go func(i int) {
			fmt.Println(i*i + 2 - 4*i)
			wg.Done()

		}(i)
	}
	wg.Wait()
}
