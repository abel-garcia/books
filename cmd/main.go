package main

import (
	"github.com/wackGarcia/books/cmd/books"
)

func main() {
	if err := books.Run(); err != nil {
		panic(err)
	}
}
