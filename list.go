package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

type ListCommand struct {
	JSON           bool `long:"json" description:"Output the list as json instead of in tabular form."`
	UnixTimestamps bool `long:"unix_time" description:"In table mode, format times as unit timestamps (seconds since the epoch)"`
}

const (
	IdColumnWidth         = 11
	TimeColumnWidthPretty = 15
	TimeColumnWidthUnix   = 10
	TypeColumnWidth       = 12

	MessageColumnMaxWidth = 40
	MessageColumnMinWidth = len("Message")
	URLColumnMaxWidth     = 30
	URLColumnMinWidth     = len("URL")
)

func truncateStr(str string, maxWidth int) string {
	if len(str) > maxWidth {
		return str[:maxWidth-3] + "..."
	}
	return str
}

func (l *ListCommand) formatTime(timestamp int64) string {
	if timestamp == 0 {
		return ""
	}

	if l.UnixTimestamps {
		return strconv.FormatInt(timestamp, 10)
	}

	t := time.Unix(timestamp, 0)
	return t.Format(time.Stamp)
}

func (l *ListCommand) Execute(args []string) error {
	checkRequiredFlags()
	postURL, err := url.Parse(options.APIHost)
	if err != nil {
		errMsg := fmt.Sprintf("Failed to parse URL %s", options.APIHost)
		return errors.New(errMsg)
	}
	postURL.Path = "/1/markers/" + options.Dataset
	req, err := http.NewRequest("GET", postURL.String(), nil)
	req.Header.Set("User-Agent", UserAgent)
	req.Header.Add("X-Honeycomb-Team", options.WriteKey)
	req.Header.Add("X-Honeycomb-Dataset", options.Dataset)
	if options.AuthorizationHeader != "" {
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", options.AuthorizationHeader))
	}
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

	if l.JSON {
		return l.ListAsJSON(body)
	} else {
		return l.ListAsTable(body)
	}
}

func (l *ListCommand) ListAsJSON(body []byte) error {
	// newlinify the JSON for one marker per line
	// TODO json-pretty-print based on a flag or something
	prettyBody := strings.Replace(string(body), "},{", "},\n{", -1)
	fmt.Println(prettyBody)
	return nil
}

func (l *ListCommand) ListAsTable(body []byte) error {
	// Unmarshal string into structs.
	var mkrs []marker
	if err := json.Unmarshal(body, &mkrs); err != nil {
		return err
	}

	urlColumnWidth := 0
	messageColumnWidth := 0

	for _, m := range mkrs {
		if len(m.Message) > messageColumnWidth {
			messageColumnWidth = len(m.Message)
		}
		if len(m.URL) > urlColumnWidth {
			urlColumnWidth = len(m.URL)
		}
	}

	if messageColumnWidth > MessageColumnMaxWidth {
		messageColumnWidth = MessageColumnMaxWidth
	}
	if messageColumnWidth < MessageColumnMinWidth {
		messageColumnWidth = MessageColumnMinWidth
	}

	if urlColumnWidth > URLColumnMaxWidth {
		urlColumnWidth = URLColumnMaxWidth
	}
	if urlColumnWidth < URLColumnMinWidth {
		urlColumnWidth = URLColumnMinWidth
	}

	var timeColumnWidth int
	if l.UnixTimestamps {
		timeColumnWidth = TimeColumnWidthUnix
	} else {
		timeColumnWidth = TimeColumnWidthPretty
	}

	fmt.Printf("| %-[2]*[1]s | %[4]*[3]s | %[6]*[5]s | %-[8]*[7]s | %-[10]*[9]s | %-[12]*[11]s |\n",
		"ID", IdColumnWidth,
		"Start Time", timeColumnWidth,
		"End Time", timeColumnWidth,
		"Type", TypeColumnWidth,
		"Message", messageColumnWidth,
		"URL", urlColumnWidth,
	)
	fmt.Printf("+-%s-+-%s-+-%s-+-%s-+-%s-+-%s-+\n",
		strings.Repeat("-", IdColumnWidth),
		strings.Repeat("-", timeColumnWidth),
		strings.Repeat("-", timeColumnWidth),
		strings.Repeat("-", TypeColumnWidth),
		strings.Repeat("-", messageColumnWidth),
		strings.Repeat("-", urlColumnWidth),
	)
	for _, m := range mkrs {
		fmt.Printf("| %-[2]*[1]s | %[4]*[3]s | %[6]*[5]s | %-[8]*[7]s | %-[10]*[9]s | %-[12]*[11]s |\n",
			m.ID, IdColumnWidth,
			l.formatTime(m.StartTime), timeColumnWidth,
			l.formatTime(m.EndTime), timeColumnWidth,
			m.Type, TypeColumnWidth,
			truncateStr(m.Message, messageColumnWidth), messageColumnWidth,
			truncateStr(m.URL, urlColumnWidth), urlColumnWidth,
		)
	}

	return nil
}
