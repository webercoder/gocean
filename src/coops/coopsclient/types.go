package coopsclient

import (
	"time"

	"github.com/webercoder/gocean/src/lib"
)

// Client interacts with the NOAA api.
type Client struct {
	Application string
	HTTPClient  lib.HTTPGetter
	URL         string
}

// ClientOption is an option type provided to the NewClient constructor.
type ClientOption func(*Client)

// ClientRequest contains data for NOAA CO-OPS API requests.
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

// ClientErrorResponse is a deserialized error response from the NOAA API.
type ClientErrorResponse struct {
	Err ClientErrorResponseMessage `json:"error"`
}

// ClientErrorResponseMessage is the error message from the ClientErrorResponse.
type ClientErrorResponseMessage struct {
	Message string `json:"message"`
}

// ClientRequestOption is an option type provided to the NewClientRequest constructor.
type ClientRequestOption func(*ClientRequest)

// Datum corresponds to the Datum pseudo-enum.
type Datum int

// Product corresponds to the Product pseudo-enum.
type Product int

// ResponseFormat corresponds to the ResponseFormat pseudo-enum.
type ResponseFormat int

// TimeZoneFormat corresponds to the TimeZoneFormat pseudo-enum.
type TimeZoneFormat int

// Units corresponds to the Units pseudo-enum.
type Units int
