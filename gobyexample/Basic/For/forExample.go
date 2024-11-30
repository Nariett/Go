package main

import (
	"fmt"
	"math"
)

func main() {
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	for j := 0; j < 3; j++ {
		fmt.Println(j)
	}

	for i := range 3 {
		fmt.Println("range", i)
	}

	for {
		fmt.Println("loop")
		break
	}

	for n := range 6 {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}

	for h := 0.1; h < 1; h += 0.1 {
		fmt.Print(math.Round(h*10)/10, " ")
	}
	fmt.Print("\n")

	// task

	for row := 1; row <= 3; row++ {
		for col := 1; col <= 3; col++ {
			fmt.Print(row*col, " ")
		}
		fmt.Print("\n")
	}
	fmt.Print("\n")

	//task palindrom
	var (
		word string
	)
	fmt.Println("Введите слово")
	fmt.Scanln(&word)
	for i := 0; i < len(word); i++ {
		if word[i] != word[len(word)-1-i] {
			fmt.Printf("Слово %s не палиндром", word)
			return
		}
	}
	fmt.Print("Слово палиндром ", word)
	//task palindrom

}
