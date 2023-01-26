# cats command

cat command with cats

```
$ cats main.go
🐈  package main
😸
😹  import (
😺  	"bufio"
😻  	"fmt"
😼  	"os"
😽  	"strings"
😾
😿  	"github.com/spf13/cobra"
🙀  )
🐈
😸  var cats = []int{0x1f408, 0x1f638, 0x1f639, 0x1f63a, 0x1f63b, 0x1f63c, 0x1f63d, 0x1f63e, 0x1f63f, 0x1f640}
😹
😺  var eFlag bool
😻  var nFlag bool
😼
😽  type options struct {
😾  	e bool
😿  	n bool
🙀  }
🐈
😸  var rootCmd = &cobra.Command{
😹  	Use:   "cat [filenames]",
😺  	Short: "Print files to the screen",
😻  	Args:  cobra.MinimumNArgs(1),
😼  	Run:   run,
😽  }
😾
...
```

## Install

```
go install github.com/nesheep/cats@latest
```
