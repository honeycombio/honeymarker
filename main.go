package main

import (
	"fmt"
	"net/http"
	"os"

	flag "github.com/jessevdk/go-flags"
)

// BuildID is set by CI
var BuildID string = "dev"

// UserAgent is what gets included in all http requests to the api
var UserAgent string

type Options struct {
	WriteKey string `short:"k" long:"writekey" env:"HONEYCOMB_API_KEY" description:"Honeycomb write key from https://ui.honeycomb.io/account"`
	Dataset  string `short:"d" long:"dataset" description:"Honeycomb dataset name from https://ui.honeycomb.io/dashboard (use __all__ for environment-wide markers)"`
	APIHost  string `long:"api_host" hidden:"true" default:"https://api.honeycomb.io/"`

	AuthorizationHeader string `long:"authorization-header" hidden:"true"`
}

var options Options
var parser = flag.NewParser(&options, flag.Default)
var client = http.Client{}
var usage = `-k <writekey> -d <dataset> COMMAND [other flags]

  honeymarker is a command line utility for manipulating markers in your
  Honeycomb dataset.

  Except for the Version command, Writekey and Dataset are both required. Most commands have additional
  arguments.

  'honeymarker COMMAND --help' will print command-specific flags`

// setVersion sets the internal version ID and updates libhoney's user-agent
func setVersionUserAgent() {
	UserAgent = fmt.Sprintf("honeymarker/%s", BuildID)
}

func checkRequiredFlags() {
	// In order to do the above, we can't use the "required" flag on the options, so we have to
	// check them manually.
	if options.WriteKey == "" {
		fmt.Println("the required flag `-k, --writekey' was not specified")
		os.Exit(1)
	}

	if options.Dataset == "" {
		fmt.Println("the required flag `-d, --dataset' was not specified")
		os.Exit(1)
	}
}

func main() {
	setVersionUserAgent()

	parser.AddCommand("add", "Add a new marker",
		`add creates a new marker with the specified attributes.

  All parameters to add are optional.

  If start_time is missing, the marker will be assigned the current time.

  It is highly recommended that you fill in either message or type.
  All markers of the same type will be shown with the same color in the UI.
  The message will be visible above an individual marker.

	If a URL is specified along with a message, the message will be shown
	as a link in the UI, and clicking it will take you to the URL.`,
		&AddCommand{})

	parser.AddCommand("list", "List all markers",
		`List all markers for the specified dataset.

  Returned markers will be displayed in tabular format by default,
	ordered by the marker's start time.`,
		&ListCommand{})

	parser.AddCommand("rm", "Delete a marker",
		`Delete the marker in the specified dataset, as identified by its ID.

	Marker IDs are available via the 'list' command.`,
		&RmCommand{})

	parser.AddCommand("update", "Update a marker",
		`Update an existing marker in the specified dataset with the specified options.

	The marker ID is required (available via the 'list' command). All other
	parameters are optional, though an 'update' will be a no-op unless a parameter
	is specified with a new value.`,
		&UpdateCommand{})

	parser.AddCommand("version", "Print tool version",
		`Prints the version number of this tool and exits.`,
		&VersionCommand{})

	// run whichever command is chosen
	parser.Usage = usage
	if _, err := parser.Parse(); err != nil {
		if flagErr, ok := err.(*flag.Error); ok {
			if flagErr.Type == flag.ErrHelp {
				// asking for help isn't a failed run.
				os.Exit(0)
			}
			if flagErr.Type == flag.ErrCommandRequired ||
				flagErr.Type == flag.ErrUnknownCommand ||
				flagErr.Type == flag.ErrUnknownFlag ||
				flagErr.Type == flag.ErrRequired {
				fmt.Println("  run 'honeymarker --help' for full usage details")
			}
		}
		os.Exit(1)
	}
}
