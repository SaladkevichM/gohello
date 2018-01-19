package main

import (
	"fmt"
	"math"
)

type Cash struct {
	Value string
}

func (c Cash) get() {
	fmt.Println(c.Value)
}

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}

	fmt.Println(v.Abs())
	fmt.Println(Abs(v))

	c := Cash{"bucks"}
	c.get()

}
