package coops

import (
	"errors"
	"flag"
	"fmt"

	"github.com/webercoder/gocean/src/coops"
)

// ResponseFormatPrettyPrint is the command-line option to pretty print results.
const ResponseFormatPrettyPrint = "pretty"

// ClientConfig holds command-line values and is used later to build a coops.ClientRequest.
type ClientConfig struct {
	BeginDate      string
	Count          int
	Datum          string
	EndDate        string
	Format         string
	Hours          int
	Station        string
	TimeZoneFormat string
	Units          string
}

// NewClientConfig creates a new ClientConfig.
func NewClientConfig() *ClientConfig {
	return &ClientConfig{}
}

// GetFlagSet returns a generic flag set for CO-OPS API usage.
func (ccc *ClientConfig) GetFlagSet(name string, errorHandling flag.ErrorHandling) *flag.FlagSet {
	fset := flag.NewFlagSet(name, errorHandling)
	fset.StringVar(&ccc.BeginDate, "begin-date", "", "The begin date for the data set.")
	fset.StringVar(
		&ccc.Datum,
		"datum",
		coops.DatumMLLW.String(),
		fmt.Sprintf("The datum to query. Possible values: %v", coops.DatumStrings),
	)
	fset.StringVar(&ccc.EndDate, "end-date", "", "The end date for the data set.")
	fset.StringVar(
		&ccc.Format,
		"format",
		ResponseFormatPrettyPrint,
		fmt.Sprintf(
			"The output format of the results. Possible values: %v",
			append(coops.ResponseFormatStrings[:], ResponseFormatPrettyPrint),
		),
	)
	fset.StringVar(&ccc.Station, "station", "", "The station to query.")
	fset.StringVar(
		&ccc.TimeZoneFormat,
		"time-zone-format",
		coops.TimeZoneFormatLSTLDT.String(),
		fmt.Sprintf("The time zone format. Possible values: %v", coops.TimeZoneFormatStrings),
	)
	fset.StringVar(
		&ccc.Units,
		"units",
		coops.UnitsEnglish.String(),
		fmt.Sprintf("Either english or metric. Possible values: %v", coops.UnitsStrings),
	)
	fset.IntVar(&ccc.Hours, "hours", 24, "The offset from the start time.")
	fset.IntVar(&ccc.Count, "count", -1, "The number of results to display. Only works with the pretty format.")
	return fset
}

// ToRequestOptions converts this config into list of coops.RequestOption's that will be used
// to build a coops.Request object.
func (ccc *ClientConfig) ToRequestOptions() ([]coops.ClientRequestOption, error) {
	datum, ok := coops.StringToDatum(ccc.Datum)
	if !ok {
		return nil, errors.New("datum param is invalid")
	}

	tzformat, ok := coops.StringToTimeZoneFormat(ccc.TimeZoneFormat)
	if !ok {
		return nil, errors.New("time-zone-format param is invalid")
	}

	units, ok := coops.StringToUnits(ccc.Units)
	if !ok {
		return nil, errors.New("units param is invalid")
	}

	format, ok := coops.StringToResponseFormat(ccc.Format)
	if !ok && ccc.Format == ResponseFormatPrettyPrint {
		format = coops.ResponseFormatJSON
	} else if !ok {
		return nil, errors.New("response format is invalid")
	}

	return []coops.ClientRequestOption{
		coops.WithBeginDateString(ccc.BeginDate),
		coops.WithDatum(datum),
		coops.WithEndDateString(ccc.EndDate),
		coops.WithFormat(format),
		coops.WithStation(ccc.Station),
		coops.WithTimeZoneFormat(tzformat),
		coops.WithUnits(units),
	}, nil
}
