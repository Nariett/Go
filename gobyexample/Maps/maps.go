package main

import (
	"fmt"
	"maps"
)

func main() {
	m := make(map[string]int)
	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map:", m)

	v1 := m["k1"]
	fmt.Println("v1:", v1)

	v3 := m["k3"]
	fmt.Println("v3:", v3)

	fmt.Println("len:", len(m))

	delete(m, "k2")
	fmt.Println("map:", m)

	clear(m)
	fmt.Println("map:", m)

	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)

	n2 := map[string]int{"foo": 1, "bar": 2}
	if maps.Equal(n, n2) {
		fmt.Println("n == n2")
	}

	//task голосование
	var (
		mapVoice   = map[string]int{"Саша": 10, "Олег": 5}
		voice      string
		maxValue   int
		winnerName string
	)

	fmt.Println("Введите имя кандидата:")
	fmt.Scanln(&voice)
	_, exists := mapVoice[voice]
	if exists {
		mapVoice[voice] += 1
	} else {
		mapVoice[voice] = 1
	}
	fmt.Printf("Голосов за %s: %d \n", voice, mapVoice[voice])

	for key, value := range mapVoice {
		if mapVoice[key] > maxValue {
			maxValue = value
			winnerName = key
		}
	}
	fmt.Printf("Победитель %s с %d голосами", winnerName, maxValue)

}
