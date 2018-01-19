package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

type Person struct {
	coord Vertex
}

var m map[string]Vertex
var m1 map[string]string

func main() {

	fmt.Println(m)

	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}

	fmt.Println(m["Bell Labs"])
	fmt.Println(m[""])

	m1 = make(map[string]string)
	m1["qwe"] = "asd"

	fmt.Println(m1["qwe"])

	p := Person{coord: Vertex{
		40.68433, -74.39967,
	}}

	fmt.Println(p)

	var pointer *Person
	pointer = &p
	fmt.Println(*pointer)

}
