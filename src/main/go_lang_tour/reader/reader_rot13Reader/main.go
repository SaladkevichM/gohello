package main

import (
	//"fmt"
	"io"
	"os"
	"strings"
)

type rot13Error struct{}

func (err rot13Error) Error() string {
	return "rot13 error"
}

type rot13Reader struct {
	reader io.Reader
}

func (rot *rot13Reader) Read(p []byte) (n int, err error) {
	n, err = rot.reader.Read(p)
	for i, value := range p {
		//fmt.Printf("%T", value)
		if (value >= 'A' && value < 'N') || (value >= 'a' && value < 'n') {
			p[i] += 13
		} else if (value > 'M' && value <= 'Z') || (value > 'm' && value <= 'z') {
			p[i] -= 13
		}
	}
	return
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{reader: s}
	io.Copy(os.Stdout, &r)
}
