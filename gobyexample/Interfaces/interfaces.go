package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perim() float64
}
type rect struct {
	width, height float64
}
type circle struct {
	radius float64
}

func (r rect) area() float64 {
	return r.width * r.height
}
func (r rect) perim() float64 {
	return 2 * r.width * 2 * r.height
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

// func measure(g geometry) {
// 	fmt.Print(g)
// 	fmt.Println(g.area())
// 	fmt.Println(g.perim())
// }

// task

type Move interface {
	move()
	getName() string
}
type Car struct {
	name string
}

func (c Car) move() {
	fmt.Println("Машина едет")
}
func (c Car) getName() string {
	return c.name
}

type Ship struct {
	name string
}

func (s Ship) move() {
	fmt.Println("Корабль плывет")
}
func (s Ship) getName() string {
	return s.name
}

func main() {
	var geoOne geometry = rect{width: 3, height: 4}
	var geoTwo geometry = circle{radius: 5}
	fmt.Println("Фигура 1", geoOne.area())
	fmt.Println(geoOne.perim())
	fmt.Println("Фигура 2", geoTwo.area())
	fmt.Println(geoTwo.perim())

	//r := rect{width: 3, height: 4}
	//c := circle{radius: 5}
	//measure(r)
	//measure(c)

	//task
	var transport Move = Car{name: "Volvo"}
	transport.move()
	fmt.Println(transport.getName())
	transport = Ship{name: "Победа"}
	transport.move()
}
