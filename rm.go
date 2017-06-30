package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type RmCommand struct {
	MarkerID string `short:"i" long:"id" description:"ID of the marker to delete" required:"true"`
}

var rmCommand RmCommand

func init() {
	parser.AddCommand("rm",
		"delete a marker",
		"delete the marker identified by ID. IDs available from the 'list' command",
		&rmCommand)
}

func (r *RmCommand) Execute(args []string) error {
	postURL, err := url.Parse(options.APIHost)
	if err != nil {
		errMsg := fmt.Sprintf("Failed to parse URL %s", options.APIHost)
		return errors.New(errMsg)
	}
	postURL.Path = "/1/markers/" + options.Dataset + "/" + r.MarkerID
	req, err := http.NewRequest("DELETE", postURL.String(), nil)
	req.Header.Add("X-Honeycomb-Team", options.WriteKey)
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		errMsg := fmt.Sprintf("Failed with %d and message: %s", resp.StatusCode, body)
		return errors.New(errMsg)
	}
	fmt.Println(string(body))
	return nil
}
