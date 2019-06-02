package main

import (
	"fmt"

	"github.com/machine23/shourl/shortener"
)

func main() {
	sh, _ := shortener.New("some.cc")

	longURL := "https://www.someurl.com/here/is/very/long/url"
	short := sh.Shorten(longURL)
	fmt.Println("Short:", short)

	// Get back the long URL
	long := sh.Resolve(short)
	fmt.Println("Long:", long)
}
