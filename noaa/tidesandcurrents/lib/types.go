package lib

import (
	"time"

	"github.com/webercoder/gocean/lib"
)

// Client interacts with the NOAA api.
type Client struct {
	Application string
	HTTPClient  lib.HTTPGetter
	URL         string
}

type ClientOption func(*Client)

type ClientRequest struct {
	BeginDate time.Time
	EndDate   time.Time
	Datum     string
	Format    string
	Product   string
	Station   string
	TimeZone  string
	Units     string
}

type ClientRequestOption func(*ClientRequest)
