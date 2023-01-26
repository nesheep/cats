package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var cats = []int{0x1f408, 0x1f638, 0x1f639, 0x1f63a, 0x1f63b, 0x1f63c, 0x1f63d, 0x1f63e, 0x1f63f, 0x1f640}

var eFlag bool
var nFlag bool

type options struct {
	e bool
	n bool
}

var rootCmd = &cobra.Command{
	Use:   "cat [filenames]",
	Short: "Print files to the screen",
	Args:  cobra.MinimumNArgs(1),
	Run:   run,
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&eFlag, "end", "e", false, "Display a dollar sign ('$') at the end of each line.")
	rootCmd.PersistentFlags().BoolVarP(&nFlag, "number", "n", false, "Number the output lines, starting at 1.")
}

func main() {
	rootCmd.Execute()
}

func run(cmd *cobra.Command, args []string) {
	opt := options{
		e: eFlag,
		n: nFlag,
	}
	for _, fn := range args {
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
