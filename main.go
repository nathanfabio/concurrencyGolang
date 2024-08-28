package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	fmt.Println(name())
	fmt.Println(surname())
	fmt.Println(city())

	fmt.Println("time spent: ", time.Since(start))
}

func name() string {
	time.Sleep(time.Millisecond * 100)
	return "nathan"
}

func surname() string {
	time.Sleep(time.Millisecond * 200)
	return "fabio"
}

func city() string {
	time.Sleep(time.Millisecond * 300)
	return "bebedouro"
}