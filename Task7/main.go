package main

import (
	"fmt"
	"sync"
	"time"
)

func writeToChan(i int, c chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	c <- i
	defer func() { <-c }()
	time.Sleep(time.Second * 20)
	fmt.Println("Обработка задачи номер: ", i)

}

func main() {
	maxTasks := 3
	countTask := 10
	ch := make(chan int, maxTasks)
	var wg sync.WaitGroup
	for i := 0; i < countTask; i++ {
		wg.Add(1)
		go writeToChan(i, ch, &wg)
	}
	wg.Wait()
	fmt.Println("Все")
}
