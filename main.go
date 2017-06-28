package main

import (
	"fmt"
	"net/http"
	"os"

	flag "github.com/jessevdk/go-flags"
)

type Options struct {
	WriteKey string `short:"k" long:"writekey" description:"Honeycomb write key from https://ui.honeycomb.io/account" required:"true"`
	Dataset  string `short:"d" long:"dataset" description:"Honeycomb dataset name from https://ui.honeycomb.io/dashboard" required:"true"`
	PostURL  string `long:"postURL" hidden:"true" default:"https://api.honeycomb.io/"`
}

var options Options
var parser = flag.NewParser(&options, flag.Default)
var client = &http.Client{}
var usage = `-k <writekey> -d <dataset> COMMAND [other flags]

  honeymarker is the command line utility for manipulating markers in your
  Honeycomb dataset.

  Writekey and Dataset are both required. Most commands have additional
  arguments.

  'honeymarker COMMAND --help' will print command-specific flags`

func main() {
	// run whichever command is chosen
	parser.Usage = usage
	if _, err := parser.Parse(); err != nil {
		if flagErr, ok := err.(*flag.Error); ok {
			if flagErr.Type == flag.ErrHelp {
				// asking for help isn't a failed run.
				os.Exit(0)
			}
			if flagErr.Type == flag.ErrCommandRequired ||
				flagErr.Type == flag.ErrUnknownFlag ||
				flagErr.Type == flag.ErrRequired {
				fmt.Println("  run 'honeymarker --help' for full usage details")
			}
		}
		os.Exit(1)
	}
}
