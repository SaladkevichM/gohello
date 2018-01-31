package main

import "fmt"

const pi = 3.14

var c, python, java bool
var i, j int = 1, 2

func add(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {

	fmt.Println(add(42, 13))

	fmt.Println(swap("world", "hello"))

	var x, y = split(2)
	fmt.Println(x)
	fmt.Println(y)

	fmt.Println(c)

	k := 3
	fmt.Println(k)
}
