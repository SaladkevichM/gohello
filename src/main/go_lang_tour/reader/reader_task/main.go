package main

import (
	"github.com/tour/reader"
)

// Реализуйте тип Reader, генерирующий бесконечный поток ASCII символа 'A'.

type MyReader struct{}

func (reader MyReader) Read(b []byte) (int, error) {

	for i := 0; i < len(b); i++ {
		b[i] = 'A'
	}

	return 1, nil
}

func main() {
	reader.Validate(MyReader{})
}
