package main

import (
	"flag"
	"fmt"
	"os"
	"sync"

	"github.com/fatih/color"
	"github.com/maxmellon/onion"
	"github.com/maxmellon/onion/formatter"
)

var style string

func init() {
	flag.StringVar(&style, "style", "simple", "Choose an output formatter.")
}

var Usage = func() {
	fmt.Fprintf(os.Stderr, "Usage of %s [options] [file1, file2, ...]", os.Args[0])
	fmt.Fprintf(os.Stderr, `
	--style, -style [FORMATTER]		choose a formatter.
		[s]imple (default)
		[e]rrorformats (vim errorformats style)
`)
	os.Exit(1)
}

func main() {
	flag.Parse()
	var argn = len(os.Args)

	if argn == 0 || argn == 1 {
		Usage()
	}

	var wg sync.WaitGroup
	for _, filename := range flag.Args() {
		wg.Add(1)
		go func(filename string) {
			data := onion.ReadFile(filename)
			parsedData := onion.Parse(data)
			lintedData := onion.Linting(parsedData)
			if style == "simple" || style == "s" {
				formatter.SimplePrintResult(filename, lintedData)
			} else if style == "errorformats" || style == "e" {
				formatter.QuickFixPrintResult(filename, lintedData)
			} else {
				styleOptParseError()
			}
			defer wg.Done()
		}(filename)
	}

	wg.Wait()
	return
}

func styleOptParseError() {
	r := color.New(color.FgRed, color.Bold)
	r.Fprintf(os.Stderr, "Unexpected format")
	y := color.New(color.FgYellow, color.Bold)
	y.Fprintf(os.Stderr, " `%s` ", style)
	r.Printf("Please choose by following formatter list")
	fmt.Fprintf(os.Stderr, `
		[s]imple (default)
		[e]rrorformats (vim errorformats style)
`)
	os.Exit(1)
}
