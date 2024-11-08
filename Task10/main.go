package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var mu sync.RWMutex
	var counter int
	for i := 0; i < 10; i++ {
		wg.Add(2)
		go func() {
			defer wg.Done()
			mu.RLock()

			x := counter
			fmt.Println(x)
			mu.RUnlock()
		}()
		go func() {
			defer wg.Done()
			mu.Lock()
			counter++
			println(counter)
			mu.Unlock()
		}()
	}
	wg.Wait()
}
