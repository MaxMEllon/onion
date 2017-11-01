package main

import (
	"fmt"
	"os"

	"github.com/maxmellon/onion"
	"github.com/maxmellon/onion/styles"
)

func main() {
	args := len(os.Args)
	if args == 0 || args == 1 {
		fmt.Printf(`
Usage:
  $ tmng-mail-lint [filename]
`)
	}

	data := onion.ReadFile(os.Args[1])
	parsedData := onion.Parse(data)
	lintedData := onion.Linting(parsedData)
	// fmt.Println(parsedData)
	// fmt.Println(lintedData)
	styles.SimplePrintResult(lintedData)
	return
}
