package main

import (
	"flag"

	"github.com/nesheep/cats"
)

func main() {
	flag.Parse()
	args := flag.Args()
	cats.Run(args)
}
