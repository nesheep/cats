package main

import (
	"flag"

	"github.com/nesheep/cats"
)

var nFlag bool

func init() {
	flag.BoolVar(&nFlag, "n", false, "Number the output lines, starting at 1.")
}

func main() {
	flag.Parse()
	args := flag.Args()
	cats.Run(args, nFlag)
}
