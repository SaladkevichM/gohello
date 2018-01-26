package main

import (
	"fmt"
	"strings"
)

func main() {

	f := "é́́é́́é́́"

	for _, v := range f {
		fmt.Println(byte(v))
	}

	s := "We went to eat at multiple café"
	cafe := "cafe"
	if p := strings.Index(s, cafe); p != -1 {
		p += len(cafe)
		s = s[:p] + "s" + s[p:]
	}
	fmt.Println(s)

}
