package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/honeycombio/hound/types"
)

type UpdateCommand struct {
	MarkerID  string `short:"i" long:"id" description:"ID for the marker to update" required:"true"`
	StartTime int64  `short:"s" long:"start_time" description:"start time for the marker in unix time"`
	EndTime   int64  `short:"e" long:"end_time" hidden:"true" description:"end time for the marker in unix time"`
	Message   string `short:"m" long:"msg" description:"message to attach to the marker"`
	URL       string `short:"u" long:"url" description:"url to attach to the marker"`
	Type      string `short:"t" long:"type" description:"type of the marker"`
}

var updateCommand UpdateCommand

func init() {
	parser.AddCommand("update",
		"update a marker",
		"update an existing marker with the specified options. IDs available from the 'list' command.",
		&updateCommand)
}

func (u *UpdateCommand) Execute(args []string) error {

	marker := types.Marker{
		StartTime: u.StartTime,
		EndTime:   u.EndTime,
		Message:   u.Message,
		Type:      u.Type,
		URL:       u.URL,
	}
	blob, err := json.Marshal(marker)
	if err != nil {
		return err
	}

	postURL, err := url.Parse(options.APIHost)
	if err != nil {
		errMsg := fmt.Sprintf("Failed to parse URL %s", options.APIHost)
		return errors.New(errMsg)
	}
	postURL.Path = "/1/markers/" + options.Dataset + "/" + u.MarkerID

	req, err := http.NewRequest("PUT", postURL.String(), bytes.NewBuffer(blob))
	req.Header.Set("Content-Type", "application/json")
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
