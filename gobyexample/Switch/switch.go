package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	i := 2
	fmt.Println("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm bool")
		case int:
			fmt.Println("I'm int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")

	//task calculator
	var (
		numberOne int
		numberTwo int
		operation string
		result    int
	)
	fmt.Println("Введите первое число: ")
	fmt.Scanf("%d\n", &numberOne)
	fmt.Println("Введите второе число: ")
	fmt.Scanf("%d\n", &numberTwo)
	fmt.Println("Выберите операцию (+, -, *, /, ^, %): ^")
	fmt.Scanf("%s\n", &operation)
	switch operation {
	case "+":
		result = numberOne + numberTwo
	case "-":
		result = numberOne - numberTwo
	case "*":
		result = numberOne * numberTwo
	case "/":
		if numberTwo == 0 {
			fmt.Println("Ошибка: деление на ноль.")
			return
		}
		result = numberOne / numberTwo
	case "^":
		result = int(math.Pow(float64(numberOne), float64(numberTwo)))
	case "%":
		if numberTwo == 0 {
			fmt.Println("Ошибка: деление на ноль.")
			return
		}
		result = numberOne % numberTwo
	default:
		fmt.Println("Ошибка: операция", operation, "не поддерживается.")
		return
	}
	fmt.Printf("Результат операции %s: %d\n", operation, result)
}
