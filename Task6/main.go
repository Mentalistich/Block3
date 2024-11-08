package main

import (
	"fmt"
	"time"
)

func longTimeExecution(channel chan<- int) {
	for i := 0; i < 10; i++ {
		time.Sleep(1 * time.Second)
		channel <- i
	}

}

func main() {

	channel := make(chan int, 10)
	timeout := time.Second * 5
	go longTimeExecution(channel)
	defer close(channel)
	for {
		select {
		case <-channel:
			fmt.Println(<-channel)
		case <-time.After(timeout):
			fmt.Println("Конец выполнения")
			return
		}
	}
}
