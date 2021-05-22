package command

import (
	"errors"
	"flag"
	"fmt"
	"os"

	"github.com/webercoder/gocean/src/coops"
)

// ResponseFormatPrettyPrint is the command-line option to pretty print results.
const ResponseFormatPrettyPrint = "pretty"

// COOPSCommandHandler is a composite of all the CO-OPS command-line commands.
type COOPSCommandHandler struct {
	subHandlers map[coops.Product]Handler
}

// NewCOOPSCommandHandler creates a new Tides and Currents CommandHandler.
func NewCOOPSCommandHandler() *COOPSCommandHandler {
	return &COOPSCommandHandler{
		subHandlers: map[coops.Product]Handler{
			coops.ProductPredictions: NewPredictionsCommandHandler(),
			coops.ProductWaterLevel:  NewWaterLevelsCommandHandler(),
		},
	}
}

// HandleCommand calls the appropriate subcommand.
func (cch *COOPSCommandHandler) HandleCommand() error {
	if len(os.Args) < 3 {
		cch.Usage(errors.New("please provide a subcommand"))
	}
	command, ok := coops.StringToProduct(os.Args[2])
	if !ok {
		cch.Usage(fmt.Errorf("unknown product: %s", command))
	}

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
	Format         string
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
