package utils

import (
	"net/http"

	"github.com/webercoder/gocean/lib"
)

// DefaultTidesEndpoint is the default tides endpoint from NOAA.
const DefaultTidesEndpoint = "https://api.tidesandcurrents.noaa.gov/api/prod/datagetter"

// Client interacts with the NOAA api.
type Client struct {
	Application string
	HTTPClient  lib.HTTPGetter
	URL         string
}

// NewClient creates a new NOAATidesClient object with default values.
func NewClient() *Client {
	return &Client{URL: DefaultTidesEndpoint, HTTPClient: &http.Client{}}
}
