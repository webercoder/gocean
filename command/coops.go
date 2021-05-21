package command

import (
	"errors"
	"flag"
	"fmt"
	"os"

	"github.com/webercoder/gocean/src/coops"
)

// COOPSCommandHandler is a composite of all the CO-OPS command-line commands.
type COOPSCommandHandler struct {
	subHandlers map[string]Handler
}

// NewCOOPSCommandHandler creates a new Tides and Currents CommandHandler.
func NewCOOPSCommandHandler() *COOPSCommandHandler {
	return &COOPSCommandHandler{
		subHandlers: map[string]Handler{
			"predictions": NewPredictionsCommandHandler(),
			"water_level": NewWaterLevelsCommandHandler(),
		},
	}
}

// HandleCommand calls the appropriate subcommand.
func (cch *COOPSCommandHandler) HandleCommand() error {
	if len(os.Args) < 3 {
		cch.Usage(errors.New("please provide a subcommand"))
	}
	command := os.Args[2]

	handler, ok := cch.subHandlers[command]
	if !ok {
		cch.Usage(fmt.Errorf("subcommand %s is not supported", command))
	}

	return handler.HandleCommand()
}

// Usage prints the usage for all subcommands of this command.
func (cch *COOPSCommandHandler) Usage(err ...error) {
	if len(err) > 0 {
		fmt.Printf("The following errors occurred: %v\n", err)
	}

	for k := range cch.subHandlers {
		fmt.Printf("  %s\n", k)
	}

	os.Exit(1)
}

// CoopsClientConfig holds command-line values and is used later to build a coops.ClientRequest.
type CoopsClientConfig struct {
	BeginDate      string
	Count          int
	Datum          string
	EndDate        string
	Hours          int
	Station        string
	TimeZoneFormat string
	Units          string
}

// NewCoopsClientConfig creates a new CoopsClientConfig.
func NewCoopsClientConfig() *CoopsClientConfig {
	return &CoopsClientConfig{}
}

// GetFlagSet returns a generic flag set for CO-OPS API usage.
func (ccc *CoopsClientConfig) GetFlagSet(name string, errorHandling flag.ErrorHandling) *flag.FlagSet {
	fset := flag.NewFlagSet("predictions", errorHandling)
	fset.StringVar(&ccc.BeginDate, "begin-date", "", "The begin date for the data set")
	fset.StringVar(&ccc.Datum, "datum", coops.DatumMLLW.String(), "The datum to query")
	fset.StringVar(&ccc.EndDate, "end-date", "", "The end date for the data set")
	fset.StringVar(&ccc.Station, "station", "", "The station to query")
	fset.StringVar(&ccc.TimeZoneFormat, "time-zone-format", coops.TimeZoneFormatLSTLDT.String(), "The time zone format")
	fset.StringVar(&ccc.Units, "units", coops.UnitsEnglish.String(), "Either english or metric")
	fset.IntVar(&ccc.Hours, "hours", 24, "The offset from the start time")
	fset.IntVar(&ccc.Count, "count", -1, "The number of results to display")
	return fset
}

// ToRequestOptions converts this config into list of coops.RequestOption's that will be used
// to build a coops.Request object.
func (ccc *CoopsClientConfig) ToRequestOptions() ([]coops.ClientRequestOption, error) {
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

	return []coops.ClientRequestOption{
		coops.WithBeginDateString(ccc.BeginDate),
		coops.WithDatum(datum),
		coops.WithEndDateString(ccc.EndDate),
		coops.WithStation(ccc.Station),
		coops.WithTimeZoneFormat(tzformat),
		coops.WithUnits(units),
	}, nil
}
