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
	StartTime int64  `short:"s" long:"start_time" description:"Start time for the marker in unix time (seconds since the epoch)"`
	EndTime   int64  `short:"e" long:"end_time" hidden:"true" description:"End time for the marker in unix time (seconds since the epoch)"`
	Message   string `short:"m" long:"msg" description:"Message describing this specific marker"`
	URL       string `short:"u" long:"url" description:"URL associated with this marker"`
	Type      string `short:"t" long:"type" description:"Identifies marker type"`
}

func (a *AddCommand) Execute(args []string) error {
	blob, err := json.Marshal(marker{
		StartTime: a.StartTime,
		EndTime:   a.EndTime,
		Message:   a.Message,
		Type:      a.Type,
		URL:       a.URL,
	})
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
