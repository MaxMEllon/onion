package main

import (
	"fmt"
	"os"

	"github.com/maxmellon/onion"
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
	fmt.Println(parsedData)
	return
}
