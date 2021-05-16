package coops_client

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
	Datum     Datum
	Format    ResponseFormat
	Product   Product
	Station   string
	TimeZone  TimeZoneFormat
	Units     Units
}

type ClientRequestOption func(*ClientRequest)

type Datum int

type Product int

type ResponseFormat int

type TimeZoneFormat int

type Units int
