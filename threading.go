package main

import (
	"fmt"
	"time"
)

func worker(workerID int, data chan int) {
	for x := range data {
		fmt.Printf("Worker %d received %d\n", workerID, x)
		time.Sleep(time.Second)
	}
}

func main() {
	canal := make(chan int)

	go worker(1, canal) // 1
	go worker(2, canal) // 2
	go worker(3, canal) // 3
	go worker(4, canal) // 4
	go worker(5, canal) // 5

	for i := 1; i <= 100; i++ {
		canal <- i
	}
}
