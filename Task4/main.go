package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}
	result := 0
	chanel := make(chan int)
	wg.Add(2)

	go func() {
		defer wg.Done()
		for {
			chanel <- result
			result = <-chanel
			fmt.Println(result)
			if result > 5 {
				break

			}

		}
	}()

	go func() {
		defer wg.Done()
		for {
			result = <-chanel
			chanel <- result + 1
			fmt.Println(result)
			if result == 5 {
				break
			}

		}
	}()

	wg.Wait()
	close(chanel)
}
