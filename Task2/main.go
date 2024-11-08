package main

import "fmt"

func main() {
	result := make(chan int)
	for i := 0; i < 10; i++ {
		go func(i int) {
			i = i * 2
			result <- i
		}(i)
		fmt.Println(<-result)
	}
	close(result)
}
