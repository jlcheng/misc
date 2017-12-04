package main

import (
	"fmt"
	"flag"
)

type parameters struct {
	url string
	verbose bool
}

func main() {
	params := parameters{}

	flag.StringVar(&params.url, "dest-url", "a", "some url or stuff")
	flag.BoolVar(&params.verbose, "verbose", false, "show verbose output")


	flag.Parse()
	fmt.Println(params)
}
