package main

import (
	"fmt"
	"math"
)

const s string = "constant"

func main() {
	fmt.Println(s)
	const number = 50000000000
	const d = 3e20 / number
	fmt.Println(d)
	fmt.Println(int64(d))
	fmt.Println(math.Sin(number))
	fmt.Println((math.Sin(number) + math.Cos(10)) / math.Pi)

	const a = 25
	const b = 30
	const mult = a * b
	fmt.Println(mult)
}
