package coopsclient

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/webercoder/gocean/src/lib"
)

// APIDateFormat is the date format used by the NOAA CO-OPS API.
const APIDateFormat = "20060102 15:04"

// NewClient returns a new CO-OPS client.
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

// Get queries the CO-OPS API for a given request.
//
// Example raw query:
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

// GetJSON queries the CO-OPS API for a given request and returns JSON data as a byte slice.
func (c *Client) GetJSON(r *ClientRequest) ([]byte, error) {
	resp, err := c.Get(r)
	if err != nil {
		return nil, fmt.Errorf("error retrieving data: %v", err)
	}

	if resp.Body != nil {
		defer resp.Body.Close()
	}

	jsonData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading predictions request body", err)
		return nil, err
	}

	return jsonData, nil
}

// WithURL overrides the default API URL for this client.
func WithURL(url string) ClientOption {
	return func(c *Client) {
		c.URL = url
	}
}

// WithHTTPClient overrides the default HTTP client for this CO-OPS client.
func WithHTTPClient(getter lib.HTTPGetter) ClientOption {
	return func(c *Client) {
		c.HTTPClient = getter
	}
}

// NewClientRequest creates a new request.
func NewClientRequest(opts ...ClientRequestOption) *ClientRequest {
	currentTime := time.Now()

	const (
		defaultDatum    = DatumMLLW
		defaultFormat   = ResponseFormatJSON
		defaultProduct  = ProductPredictions
		defaultTimeZone = TimeZoneFormatLSTLDT
		defaultUnits    = UnitsEnglish

		// Used to calculate time range
		defaultHours = 12
	)

	r := &ClientRequest{
		BeginDate: currentTime,
		Datum:     defaultDatum,
		Format:    defaultFormat,
		Product:   defaultProduct,
		TimeZone:  defaultTimeZone,
		Units:     defaultUnits,
	}

	for _, opt := range opts {
		opt(r)
	}

	if r.EndDate.IsZero() {
		WithHours(defaultHours)(r)
	}

	return r
}

// GetURLValues retrieves the settings in this request as a url.Values structure.
func (r *ClientRequest) GetURLValues() *url.Values {
	params := &url.Values{}
	params.Add("begin_date", r.BeginDate.Format("20060102 15:04"))
	params.Add("datum", r.Datum.String())
	params.Add("end_date", r.EndDate.Format("20060102 15:04"))
	params.Add("format", r.Format.String())
	params.Add("product", r.Product.String())
	params.Add("station", r.Station)
	params.Add("time_zone", r.TimeZone.String())
	params.Add("units", r.Units.String())
	return params
}

// WithBeginDate overrides the default being date for the result set.
func WithBeginDate(d time.Time) ClientRequestOption {
	return func(r *ClientRequest) {
		r.BeginDate = d
	}
}

// WithDatum overrides the default NOAA tide datum.
func WithDatum(datum Datum) ClientRequestOption {
	return func(r *ClientRequest) {
		r.Datum = datum
	}
}

// WithEndDate overrides the default end date for the result set.
func WithEndDate(d time.Time) ClientRequestOption {
	return func(r *ClientRequest) {
		r.EndDate = d
	}
}

// WithFormat overrides the default response format.
func WithFormat(format ResponseFormat) ClientRequestOption {
	return func(r *ClientRequest) {
		r.Format = format
	}
}

// WithProduct sets product to query.
func WithProduct(product Product) ClientRequestOption {
	return func(r *ClientRequest) {
		r.Product = product
	}
}

// WithStation sets NOAA measuring station to query.
func WithStation(station string) ClientRequestOption {
	return func(r *ClientRequest) {
		r.Station = station
	}
}

// WithTimeZoneFormat changes the time zone format of the request.
func WithTimeZoneFormat(tzf TimeZoneFormat) ClientRequestOption {
	return func(r *ClientRequest) {
		r.TimeZone = tzf
	}
}

// WithUnits sets the measuring units for the response data.
func WithUnits(units Units) ClientRequestOption {
	return func(r *ClientRequest) {
		r.Units = units
	}
}

// WithHours is a convenience function for setting the number of hours in the future
// from the BeginDate. Not to be confused with the range parameter, which is not currently
// supported by this library.
func WithHours(hours int) ClientRequestOption {
	return func(r *ClientRequest) {
		r.EndDate = r.BeginDate.Add(time.Hour * time.Duration(hours))
	}
}
