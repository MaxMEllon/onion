package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/TomiLabo/onion"
	"github.com/TomiLabo/onion/formatter"
	"github.com/TomiLabo/tmngparser"
	"github.com/TomiLabo/tmngparser/builder"
	"github.com/fatih/color"
)

var (
	style   string
	version bool
	fix     bool
)

const VERSION = `0.1.0`

func init() {
	flag.StringVar(&style, "f", "simple", "Choose an output formatter.")
	flag.BoolVar(&version, "v", false, "version")
	flag.BoolVar(&fix, "fix", false, "version")
}

var Usage = func() {
	fmt.Fprintf(os.Stderr, "Usage of %s [options] [file1, file2, ...]", os.Args[0])
	fmt.Fprintf(os.Stderr, `
	-f=FORMATTER]		choose a formatter.
		[s]imple (default)
		[e]rrorformats (vim errorformats style)
`)
	fmt.Fprintf(os.Stderr, `
	--fix -fix			auto fixing.
`)
	os.Exit(1)
}

func main() {
	flag.Parse()
	var argn = len(os.Args)

	if argn == 0 || argn == 1 {
		Usage()
	}

	if version == true {
		fmt.Println(VERSION)
		os.Exit(0)
	}

	if fix == true {
		for _, filename := range flag.Args() {
			data := parser.ReadFile(filename)
			parsed := parser.Parse(data)
			fixed := onion.Fixing(parsed)
			content := builder.BuildFromTree(fixed)
			ioutil.WriteFile(filename, []byte(content), os.ModePerm)
		}
	}

	for _, filename := range flag.Args() {
		data := parser.ReadFile(filename)
		parsedData := parser.Parse(data)
		lintedData := onion.Linting(parsedData)
		if style == "simple" || style == "s" {
			formatter.SimplePrintResult(filename, lintedData)
		} else if style == "errorformats" || style == "e" {
			formatter.QuickFixPrintResult(filename, lintedData)
		} else {
			failParseFormatterOpt()
		}
	}

	return
}

func failParseFormatterOpt() {
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
