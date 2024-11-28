package main

import (
	"fmt"
	"math"
	"strings"
)

func plus(a int, b int) int {
	return a + b
}

func plusPlus(a, b, c int) int {
	return a + b + c
}

func showResult(a, b int) {
	fmt.Println("Result:", plus(a, b))
}

// task сортировка по длине строки
func sortStringsByLength(strings []string) []string {
	length := len(strings)
	for i := 0; i < length-1; i++ {
		for j := 0; j < length-1-i; j++ {
			if len(strings[j]) > len(strings[j+1]) {
				temp := strings[j]
				strings[j] = strings[j+1]
				strings[j+1] = temp
			}
		}
	}
	return strings
}

// task поиск НОД
func gcd(a, b int) int {
	for i := a; i > 0; i-- {
		if a%i == 0 && b%i == 0 {
			return i
		}
	}
	return 1
}

// task реверс строки
func reverseString(s string) string {
	var arr []string
	for i := len(s) - 1; i >= 0; i-- {
		arr = append(arr, string(s[i]))
	}
	return strings.Join(arr, "")
}

// task проверка числа на простоту
func isPrime(number int) bool {
	if number <= 1 {
		return false
	}
	for i := 2; i <= int(math.Sqrt(float64(number))); i++ {
		if number%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	res := plus(1, 2)
	fmt.Println("1+2 =", res)

	res = plusPlus(1, 2, 3)
	fmt.Println("1+2+3 =", res)

	showResult(1, 2)

	arr := []string{"apple", "banana", "kiwi", "strawberry"}
	fmt.Println("before sort:", arr)

	arr = sortStringsByLength(arr)
	fmt.Println("after sort:", arr)

	fmt.Println(gcd(12, 18))
	number := 6
	if isPrime(number) {
		fmt.Println(number, "— простое число")
	} else {
		fmt.Println(number, "не является простым числом")
	}
	fmt.Println(reverseString("hello world"))
}
