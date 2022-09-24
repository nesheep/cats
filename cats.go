package cats

import (
	"bufio"
	"fmt"
	"os"
)

var cats = []int{0x1f408, 0x1f638, 0x1f639, 0x1f63a, 0x1f63b, 0x1f63c, 0x1f63d, 0x1f63e, 0x1f63f, 0x1f640}

func Run(fileNames []string, withNumber bool) {
	for _, fn := range fileNames {
		err := printFile(fn, withNumber)
		if err != nil {
			continue
		}
	}
}

func printFile(filename string, withNumber bool) error {
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
		if withNumber {
			line = fmt.Sprintf("%6d  %s", i+1, line)
		}
		fmt.Print(line)
		i++
	}

	return nil
}
