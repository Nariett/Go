package main

import "fmt"

func fact(num int) int {
	if num == 0 {
		return 1
	}
	return num * fact(num-1)
}
func main() {
	fmt.Println(fact(7))

	var fib func(n int) int

	fib = func(n int) int {
		if n < 2 {
			return n
		}
		return fib(n-1) + fib(n-2)
	}
	fmt.Print(fib(7))
}
