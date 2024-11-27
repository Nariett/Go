package main

import "fmt"

func main() {
	var a [5]int
	fmt.Println("emp: ", a)
	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	fmt.Println("len: ", len(a))

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl one: ", b)

	b = [...]int{1, 2, 3, 4, 5}
	fmt.Println("dcl two: ", b)

	b = [...]int{100, 2, 3, 4, 5}
	fmt.Println("idx: ", b)

	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

	twoD = [2][3]int{
		{1, 3, 2},
		{4, 3, 6},
	}
	fmt.Println("2d: ", twoD)

	//task поиск уникальных значений
	var (
		arr     = []int{1, 2, 3, 2, 4, 5, 0, 6}
		arrCont []int
		flag    int = -1
	)
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr); j++ {
			if arr[i] == arr[j] {
				flag++
			}
		}
		if flag == 0 {
			arrCont = append(arrCont, arr[i])
			flag = -1
		}
		flag = -1

	}
	fmt.Println(arrCont)

	//task второе по величине значение
	var (
		array           = []int{1, 2, 3, 4, 5}
		maxValue    int = array[0]
		preMaxValue int = -1
	)
	fmt.Println("Массив: ", array)
	for i := 0; i < len(array); i++ {
		if array[i] > maxValue {
			preMaxValue = maxValue
			maxValue = array[i]
		} else if array[i] > preMaxValue && array[i] != maxValue {
			preMaxValue = array[i]
		}
	}
	if maxValue == preMaxValue {
		fmt.Println("Ошибка: Второе по величине число не найдено.")
	} else {
		fmt.Println("Второе по значению: ", preMaxValue)
	}
}
