package main

import "fmt"

func main() {
	result := make(chan int)
	defer close(result)
	for i := 0; i < 5; i++ {
		go func(i int) {
			result <- i

		}(i)
		fmt.Println(<-result)
	}
	//for i := 0; i < 5; i++ {
	//	fmt.Println(<-result)
}
