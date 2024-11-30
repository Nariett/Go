package main

import "fmt"

type base struct {
	num int
}

func (b base) describe() string {
	return fmt.Sprintf("base with num=%v", b.num)
}

type container struct {
	base
	str string
}

// task Car
type engine struct {
	name   string
	volume int
}

func (e engine) String() string {
	return fmt.Sprintf("Название двигателя: %s\nОбъем двигателя: %d", e.name, e.volume)
}

type body struct {
	typeBody string
	seat     int
}

func (b body) String() string {
	return fmt.Sprintf("Тип кузова: %s\nКол-во мест: %d", b.typeBody, b.seat)
}

type car struct {
	carName string
	engine
	body
}

func (c car) String() string {
	return fmt.Sprintf("Название автомобиля: %s\n%s\n%s", c.carName, c.engine.String(), c.body.String())
}

func main() {
	co := container{
		base: base{
			num: 1,
		},
		str: "some one",
	}
	fmt.Printf("co={num: %v, str: %v}\n", co.num, co.str)

	fmt.Println("also num:", co.base.num)

	fmt.Println("describe:", co.describe())

	type describe interface {
		describe() string
	}
	var d describe = co
	fmt.Println("describer:", d.describe())

	//task car

	car := car{
		carName: "Toyota",
		engine: engine{
			name:   "2.0 GT",
			volume: 2,
		},
		body: body{
			typeBody: "Универсал",
			seat:     6,
		},
	}
	fmt.Println(car)
}
