package main

import "fmt"

func main() {

	primes := [6]int{2, 3, 5, 7, 11, 13}

	var slice []int = primes[1:4]

	fmt.Println(primes)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))

	slice[0] = 111

	fmt.Println(primes)

	fmt.Println("==================")

	s := []int{2, 3, 5, 7, 11, 13}

	s = s[1:4]
	fmt.Println(s)

	s = s[:2]
	fmt.Println(s)

	s = s[1:]
	fmt.Println(s)

	fmt.Println("==================")

	var a [10]int

	// there are some ways to create same slices below
	a1 := a[0:10]
	a2 := a[:10]
	a3 := a[0:]
	a4 := a[:]

	fmt.Println(a1, a2, a3, a4)

}
