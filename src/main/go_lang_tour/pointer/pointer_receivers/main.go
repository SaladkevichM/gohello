package main

/*
функции с аргументом указателем должны принимать указатель:

var v Vertex
ScaleFunc(v)  // Ошибка компиляции!
ScaleFunc(&v) // OK
тогда как методы объявленные с получателем-указателем могут принимать как значение, так и указатель при вызове:

var v Vertex
v.Scale(5)  // OK
p := &v
p.Scale(10) // OK
*/

import (
	"fmt"
)

type Vertex struct {
	X, Y float64
}

func (v *Vertex) printMethod() {
	fmt.Println(v.X, v.Y)
}

func printFunc(v *Vertex) {
	fmt.Println(v.X, v.Y)
}

func main() {

	v := Vertex{1, 2}
	v.printMethod() // Go в качестве удобства интерпретирует v.printMethod() как (&v).printMethod()
	(&v).printMethod()

	p := &v

	printFunc(p)
	p.printMethod() // вызов метода p.printMethod() интерпретируется как (*p).printMethod().
	(*p).printMethod()

	// ===============

}
