package main

import "fmt"

/*

Функция make создает обнуленный массив и возвращает срез, который ссылается на этот массив.

a := make([]int, 5)  // len(a)=5
Чтобы указать вместимость, укажите третий аргумент к make:

b := make([]int, 0, 5) // len(b)=0, cap(b)=5

b = b[:cap(b)] // len(b)=5, cap(b)=5
b = b[1:]      // len(b)=4, cap(b)=4

*/

func main() {
	a := make([]int, 5, 10) // capacity = 10, len = 5
	a = a[:10]              // increase slice capacity to 10
	a[7] = 1
	printSlice("a", a)

	b := make([]int, 0, 5)
	printSlice("b", b)

	c := b[:2]
	printSlice("c", c)

	d := c[2:5]
	printSlice("d", d)
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}
