package main

import "fmt"

func main() {
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}

	if 8%2 == 0 || 7%2 == 0 {
		fmt.Println("either 8 or 7 are even")
	}

	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}

	//task

	var (
		starnNumber  int
		finishNumber int
	)
	fmt.Println("Введите диапазон чисел\nВведите начально число: ")
	fmt.Scanf("%d\n", &starnNumber)
	fmt.Println("Введите конечное число: ")
	fmt.Scanf("%d\n", &finishNumber)
	for i := starnNumber; i < finishNumber; i++ {
		if i%2 == 0 {
			fmt.Println(i, "- четное")
		} else {
			fmt.Println(i, "- нечетное")
		}
	}

	// task 2 ромб
	var (
		size      int
		startSize int = 1
		spaceSize int
		flag      bool = true
	)
	fmt.Println("Введите размер ромба")
	fmt.Scanf("%d/n", &size)
	spaceSize = size / 2
	for i := 0; i < size; i++ {
		if spaceSize == 0 {
			flag = false
		}
		for j := 0; j < spaceSize; j++ {
			fmt.Print(" ")
		}
		for z := 0; z < startSize; z++ {
			fmt.Print("*")
		}
		if flag {
			startSize += 2
			spaceSize--
		} else {
			startSize -= 2
			spaceSize++
		}
		fmt.Println()
	}
}
