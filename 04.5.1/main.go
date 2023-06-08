package main

import (
	"fmt"
	"math"
)

func main() {
	r := Rectangle{
		width:  5,
		height: 7,
	}
	r.Type()
	r.Area()

	c := Circle{
		r: 3,
	}
	c.Type()
	c.Area()
}

type Rectangle struct {
	width  int
	height int
}

func (r *Rectangle) Area() {
	fmt.Println(r.width * r.height)
}

func (r *Rectangle) Type() {
	fmt.Println("Прямоугольник")
}

type Circle struct {
	r float64
}

func (c *Circle) Area() {
	fmt.Println(math.Pi * c.r * c.r)
}

func (c *Circle) Type() {
	fmt.Println("Окружность")
}
