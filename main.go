package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	start := time.Now()
	var wg sync.WaitGroup
	wg.Add(3)
	channel := make(chan string, 3)

	go name(&wg, channel)
	go surname(&wg, channel)
	go city(&wg, channel)
	wg.Wait()

	close(channel)

	for value := range channel {
		fmt.Println(value)
	}

	fmt.Println("time spent: ", time.Since(start))
}

func name(wg *sync.WaitGroup, channel chan string) {
	defer wg.Done()
	time.Sleep(time.Millisecond * 100)
	channel <- "nathan"
}

func surname(wg *sync.WaitGroup, channel chan string) {
	defer wg.Done()
	time.Sleep(time.Millisecond * 200)
	channel <- "fabio"
}

func city(wg *sync.WaitGroup, channel chan string) {
	defer wg.Done()
	time.Sleep(time.Millisecond * 300)
	channel <- "bebedouro"
}