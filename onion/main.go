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

	for _, filename := range os.Args[1:] {
		data := onion.ReadFile(filename)
		parsedData := onion.Parse(data)
		lintedData := onion.Linting(parsedData)
		// styles.SimplePrintResult(filename, lintedData)
		styles.QuickFixPrintResult(filename, lintedData)
	}
	return
}
