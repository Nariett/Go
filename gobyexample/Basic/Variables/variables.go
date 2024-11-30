package main

import "fmt"

func main() {
	var (
		name   string  = "Sasha"
		age            = 20
		height float32 = 172.12
	)

	fmt.Println("My name is", name, "\nI'm", age, "years old\nMy height is", height, "cm")

	var a = "initial"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	var e int
	fmt.Println(e)

	f := "apple"
	fmt.Println(f)
}
