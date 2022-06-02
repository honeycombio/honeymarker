package main

import (
	"fmt"
	"os"
)

type VersionCommand struct {
}

// Prints the version number and exits
func (v *VersionCommand) Execute(args []string) error {
	fmt.Println(BuildID)
	os.Exit(0)
	return nil
}
