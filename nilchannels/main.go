package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	a := asChan(1, 3, 5, 7)
	for v := range a {
		fmt.Println(v)
	}
}

func asChan(vs ...int) <-chan int {
	c := make(chan int)
	go func() {
		for _, v := range vs {
			c <- v
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		}
		close(c)
	}()
	return c
}
