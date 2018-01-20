package main

import "fmt"

func main() {

	primes := [3]int{2, 3, 5}

	p := &primes
	fmt.Println(*p)
	p[0] = 0
	fmt.Println(*p)

	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64) // won't panic, ok contains assertion success
	fmt.Println(f, ok)

	//f = i.(float64) // panic
	//fmt.Println(f)

}
