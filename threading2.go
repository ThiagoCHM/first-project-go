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

	qtdWorkers := 50

	for i := 1; i < qtdWorkers; i++ {
		go worker(i, canal)
	}

	go worker(1, canal)

	for i := 1; i <= 100; i++ {
		canal <- i
	}
}
