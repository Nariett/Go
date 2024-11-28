package main

import (
	"fmt"
	"slices"
)

func main() {
	var s []string
	fmt.Println("uninit: ", s, s == nil, len(s) == 0)

	s = make([]string, 3)
	fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])
	fmt.Println("len:", len(s))
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	l := s[2:5]
	fmt.Println("sl1:", l)

	l = s[:5]
	fmt.Println("sl2:", l)

	l = s[2:]
	fmt.Println("sl3:", l)

	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	t2 := []string{"g", "h", "i"}
	if slices.Equal(t, t2) {
		fmt.Println("t == t2")
	}

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d:", twoD)

	//task библиотека

	var (
		books []string
		flag  bool = false
	)
	for {
		var (
			value    int
			bookName string
		)

		fmt.Println("Выберите действие:\n1 - добавить книгу\n2 - удалить книгу\n3 - показать все книги в библиотеке\n4 - найти книгу\n5 - выйти из программы")
		fmt.Scanf("%d\n", &value)
		switch value {
		case 1:
			fmt.Println("Введите название книги: ")
			fmt.Scan(&bookName)
			books = append(books, bookName)
			fmt.Printf("Книга %s была успешно добавлена в библиотеку", bookName)
			fmt.Println("\nСписок ваших книг:", books)

		case 2:
			fmt.Println("Введите название книги: ")
			fmt.Scan(&bookName)
			if slices.Contains(books, bookName) {
				index := slices.Index(books, bookName)
				books = append(books[:index], books[index+1:]...)
				fmt.Println("Список ваших книг:", books)
			} else {
				fmt.Printf("Книга %s не была найдена в библиотеке", bookName)
			}
		case 3:
			fmt.Println("Список ваших книг:", books)
		case 4:
			fmt.Println("Введите название книги, которую желаете найти: ")
			fmt.Scan(&bookName)
			if slices.Contains(books, bookName) {
				fmt.Printf("Книга %s была успешно найдена в библиотеке", bookName)
				fmt.Println()
			} else {
				fmt.Printf("Книга %s не была найдена в библиотеке", bookName)
			}
		case 5:
			flag = true
			bookName = ""
		}
		if flag {
			break
		}
	}

}
