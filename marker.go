package main

import "time"

// The marker type, as described by https://honeycomb.io/docs/reference/api/#markers
type marker struct {
	ID string `json:"id"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	// StartTime unix timestamp truncates to seconds
	StartTime int64 `json:"start_time,omitempty"`
	// EndTime unix timestamp truncates to seconds
	EndTime int64 `json:"end_time,omitempty"`
	// Message is optional free-form text associated with the message
	Message string `json:"message,omitempty"`
	// Type is an optional marker identifier, eg 'deploy' or 'chef-run'
	Type string `json:"type,omitempty"`
	// URL is an optional url associated with the marker
	URL string `json:"url,omitempty"`
	// Color is not stored in the marker table but populated by a join
	Color string `json:"color,omitempty"`
}
