package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{}
	v.X = 4
	fmt.Println(v.X)
	fmt.Println(v.Y)

	p := &v
	p.X = 111
	fmt.Println(v)

	v2 := Vertex{X: 123} // Y:0 is implicit
	fmt.Println(v2.X)

}
