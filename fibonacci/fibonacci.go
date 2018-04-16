package main

import "fmt"

func Fib(n int) int {
	switch n {
	case 0:
		return 0
	case 1:
		return 1
	default:
		return Fib(n-1) + Fib(n-2)
	}
}

func main() {
	f := Fib(10)
	fmt.Println(f)
}
