package lib

import (
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/webercoder/gocean/lib"
)

const APIDateFormat = "20060102 15:04"

func NewClient(application string, opts ...ClientOption) *Client {
	const defaultTidesEndpoint = "https://api.tidesandcurrents.noaa.gov/api/prod/datagetter"

	c := &Client{
		Application: application,
		URL:         defaultTidesEndpoint,
	}

	for _, opt := range opts {
		opt(c)
	}

	if c.HTTPClient == nil {
		c.HTTPClient = &http.Client{}
	}

	return c
}

// Example query:
// https://api.tidesandcurrents.noaa.gov/api/prod/datagetter?station=9410170&
// product=predictions&units=metric&time_zone=lst_ldt&application=gocean&format=json&
// datum=STND&begin_date=20210119&end_date=20210121
func (c *Client) Get(r *ClientRequest) (*http.Response, error) {
	baseURL, err := url.Parse(c.URL)
	if err != nil {
		fmt.Println("Unable to parse API URL", err)
		return nil, fmt.Errorf("unable to parse API URL: %v", err)
	}

	params := r.GetURLValues()
	params.Add("application", c.Application)
	baseURL.RawQuery = params.Encode()

	return c.HTTPClient.Get(baseURL.String())
}

func WithURL(url string) ClientOption {
	return func(c *Client) {
		c.URL = url
	}
}

func WithHTTPClient(getter lib.HTTPGetter) ClientOption {
	return func(c *Client) {
		c.HTTPClient = getter
	}
}

func NewClientRequest(opts ...ClientRequestOption) *ClientRequest {
	currentTime := time.Now()

	const (
		defaultDatum    = "MLLW"
		defaultFormat   = "json"
		defaultProduct  = "predictions"
		defaultTimeZone = "lst_ldt"
		defaultUnits    = "english"

		// Used to calculate time range
		defaultHours = 12
	)

	r := &ClientRequest{
		Datum:    defaultDatum,
		Format:   defaultFormat,
		Product:  defaultProduct,
		TimeZone: defaultTimeZone,
		Units:    defaultUnits,
	}

	for _, opt := range opts {
		opt(r)
	}

	if r.BeginDate.IsZero() {
		r.BeginDate = currentTime
	}

	if r.EndDate.IsZero() {
		r.EndDate = currentTime.Add(time.Hour * time.Duration(defaultHours))
	}

	return r
}

func (r *ClientRequest) GetURLValues() *url.Values {
	params := &url.Values{}
	params.Add("begin_date", r.BeginDate.Format("20060102 15:04"))
	params.Add("end_date", r.EndDate.Format("20060102 15:04"))
	params.Add("datum", r.Datum)
	params.Add("format", r.Format)
	params.Add("product", r.Product)
	params.Add("station", r.Station)
	params.Add("time_zone", r.TimeZone)
	params.Add("units", r.Units)
	return params
}

func WithStation(station string) ClientRequestOption {
	return func(r *ClientRequest) {
		r.Station = station
	}
}