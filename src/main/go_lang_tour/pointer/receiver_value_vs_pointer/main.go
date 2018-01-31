package main

import (
	"fmt"
	"math"
)

/*
Есть две причины, чтобы использовать указатель для получателя.

Во-первых, метод может изменять значение, на которое ссылается указатель.

Во-вторых, тем самым можно избежать копирования значения при каждом вызове метода. Это может быть более эффективно, если получателем является большая структура, например.

В данном примере у Scale и Abs в качестве получателя указан тип *Vertex, хотя метод Abs не имеет необходимости изменять получателя.

В общем, все методы для типа должны иметь либо значение, либо указатель в качестве получателя, но не оба сразу.
*/

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v Vertex) ScaleX(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := &Vertex{3, 4}
	fmt.Printf("Before scaling: %+v, Abs: %v\n", v, v.Abs())
	v.Scale(5)
	fmt.Printf("After scaling: %+v, Abs: %v\n", v, v.Abs())
}
