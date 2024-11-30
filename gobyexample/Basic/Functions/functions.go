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

func vals() (int, int) {
	return 3, 7
}

func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
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

// task возращение ошибки и сообщения
func errorMessage() (int, string) {
	return 404, "Страница не найдена"
}

// task соединить строк
func connetcString(str ...string) {
	fmt.Print(str, " ")
	var result string
	for _, s := range str {
		result += s + " "
	}
	fmt.Println(result)
}

func main() {
	// Example functions
	res := plus(1, 2)
	fmt.Println("1+2 =", res)

	res = plusPlus(1, 2, 3)
	fmt.Println("1+2+3 =", res)

	//Example Multiple Return Values
	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	_, c := vals()
	fmt.Println(c)

	//Example Variadic Functions
	sum(1, 2)
	sum(1, 2, 3)

	nums := []int{1, 2, 3, 4}
	sum(nums...)

	//task

	fmt.Println("Task")
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

	code, message := errorMessage()
	fmt.Printf("Ошибка: %d\nТекст ошибики: %s\n", code, message)

	connetcString("Привет", "Саша")

}
