package main

import (
	"fmt"

	"github.com/jbbarquero/gobenchmarking/fibonacci"
)

func main() {
	f := fibonacci.Fib(10)
	fmt.Println(f)
}
