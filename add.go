package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type AddCommand struct {
	StartTime int64  `short:"s" long:"start_time" description:"start time for the marker in unix time (seconds since the epoch)"`
	EndTime   int64  `short:"e" long:"end_time" hidden:"true" description:"end time for the marker in unix time (seconds since the epoch)"`
	Message   string `short:"m" long:"msg" description:"message describing this specific marker"`
	URL       string `short:"u" long:"url" description:"URL associated with this marker"`
	Type      string `short:"t" long:"type" description:"identifies marker type"`
}

var addCommand AddCommand

const add_usage = `add creates a new marker with the specified attributes

  All parameters to add are optional.

  If start_time is missing, the marker will be assigned the current time.

  It is highly recommended that you fill in either message or type.
  All markers of the same type will be shown with the same color in the UI.
  The message will be visible above an individual marker.

	If a URL is specified along with a message, the message will be shown
	as a link in the UI, and clicking it will take you to the URL.`

func init() {
	parser.AddCommand("add",
		"add a new marker",
		add_usage,
		&addCommand)
}

func (a *AddCommand) Execute(args []string) error {

	marker := Marker{
		StartTime: a.StartTime,
		EndTime:   a.EndTime,
		Message:   a.Message,
		Type:      a.Type,
		URL:       a.URL,
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
	postURL.Path = "/1/markers/" + options.Dataset
	req, err := http.NewRequest("POST", postURL.String(), bytes.NewBuffer(blob))
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
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		errMsg := fmt.Sprintf("Failed with %d and message: %s", resp.StatusCode, body)
		return errors.New(errMsg)
	}
	fmt.Println(string(body))
	return nil

}
