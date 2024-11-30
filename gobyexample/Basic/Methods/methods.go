package main

import "fmt"

type rect struct {
	width, height int
}

func (r *rect) area() int {
	return r.width * r.height
}

func (r *rect) xd() {
	r.height = 999
}

func makeRect(w, h int) *rect {
	rect := rect{width: w, height: h}
	return &rect
}

func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func main() {
	r := rect{width: 10, height: 5}
	fmt.Println("area: ", r.area())
	fmt.Println("perim:", r.perim())
	r.xd()
	fmt.Println(r)

	rect := makeRect(10, 15)
	fmt.Println(rect)

}
