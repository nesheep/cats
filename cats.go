package cats

import (
	"bufio"
	"fmt"
	"os"
)

func Run(fileNames []string) {
	for _, fn := range fileNames {
		err := printFile(fn)
		if err != nil {
			continue
		}
	}
}

func printFile(filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	return nil
}
