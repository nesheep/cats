package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

var cats = []int{0x1f408, 0x1f638, 0x1f639, 0x1f63a, 0x1f63b, 0x1f63c, 0x1f63d, 0x1f63e, 0x1f63f, 0x1f640}

var eFlag bool
var nFlag bool

type options struct {
	e bool
	n bool
}

func init() {
	flag.BoolVar(&eFlag, "e", false, "Display a dollar sign ('$') at the end of each line.")
	flag.BoolVar(&nFlag, "n", false, "Number the output lines, starting at 1.")
}

func main() {
	flag.Parse()
	args := flag.Args()
	opt := options{
		e: eFlag,
		n: nFlag,
	}
	run(args, opt)
}

func run(fileNames []string, opt options) {
	for _, fn := range fileNames {
		err := printFile(fn, opt)
		if err != nil {
			continue
		}
	}
}

func printFile(filename string, opt options) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	i := 0
	for s.Scan() {
		cat := cats[(i)%len(cats)]
		line := fmt.Sprintf("%c  %s\n", cat, s.Text())
		if opt.e {
			line = strings.Replace(line, "\n", "$\n", 1)
		}
		if opt.n {
			line = fmt.Sprintf("%6d  %s", i+1, line)
		}
		fmt.Print(line)
		i++
	}

	return nil
}
