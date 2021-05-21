# gocean

[![<ORG_NAME>](https://circleci.com/gh/webercoder/gocean.svg?style=svg)](https://circleci.com/gh/webercoder/gocean)
[![Go Report Card](https://goreportcard.com/badge/github.com/webercoder/gocean)](https://goreportcard.com/report/github.com/webercoder/gocean)
[![Maintainability](https://api.codeclimate.com/v1/badges/f9d7f2157e1538a06b13/maintainability)](https://codeclimate.com/github/webercoder/gocean/maintainability)

gocean is a set of NOAA API wrappers and tools written in Go. It is currently under initial development and not ready for external consumption.

## Installation

To use the library:

```txt
go get -u github.com/webercoder/gocean
```

To install the binary:

```txt
go install github.com/webercoder/gocean
```

## Command-Line Usage

### Get the Nearest Station

```txt
gocean stations
```

Supported parameters ([using Go flag syntax](https://golang.org/pkg/flag/#hdr-Command_line_flag_syntax)):

```txt
-postcode string
    Find stations near this postcode
```

Example:

```txt
$ gocean stations -postcode 94087
Finding nearest station to 94087
The nearest Station is "Redwood City" (ID: 9414523), which is 23.072995 kms away from 94087.

$ gocean stations -postcode 92101
Finding nearest station to 92101
The nearest Station is "San Diego, San Diego Bay" (ID: 9410170), which is 1.130777 kms away from 92101.
```

### Get Tide Predictions

```txt
gocean coops predictions
```

Supported parameters ([using Go flag syntax](https://golang.org/pkg/flag/#hdr-Command_line_flag_syntax)):

```txt
-begin-date string
    The begin date for the data set
-count int
    The number of results to display (default -1)
-datum string
    The datum to query (default "MLLW")
-end-date string
    The end date for the data set
-hours int
    The offset from the start time (default 24)
-station string
    The station to query
-time-zone-format string
    The time zone format (default "lst_ldt")
-units string
    Either english or metric (default "english")
```

Example:

```txt
$ gocean coops predictions -station 9414523 -count 5
Tide predictions for station: 9414523
  2021-05-21 07:18  6.106
  2021-05-21 07:24  6.172
  2021-05-21 07:30  6.232
  2021-05-21 07:36  6.287
  2021-05-21 07:42  6.337
```

### Get Water Levels

```txt
gocean coops water_level
```

Supported parameters ([using Go flag syntax](https://golang.org/pkg/flag/#hdr-Command_line_flag_syntax)):

```txt
-begin-date string
    The begin date for the data set
-count int
    The number of results to display (default -1)
-datum string
    The datum to query (default "MLLW")
-end-date string
    The end date for the data set
-hours int
    The offset from the start time (default 24)
-station string
    The station to query
-time-zone-format string
    The time zone format (default "lst_ldt")
-units string
    Either english or metric (default "english")
```

Example:

```txt
$ gocean coops water_level -station 9410230 -count 5
Tide water levels for station: 9410230
  2021-05-20 07:24  2.734  Preliminary
  2021-05-20 07:30  2.678  Preliminary
  2021-05-20 07:36  2.556  Preliminary
  2021-05-20 07:42  2.458  Preliminary
  2021-05-20 07:48  2.376  Preliminary
```

## Code Structure

Directory | Description
--- | ---
`/command` | Command-line options for the main program.
`/src/coops` | Libraries for connecting to the [NOAA CO-OPS API](https://api.tidesandcurrents.noaa.gov/api/prod/).
`/src/stations` | Libraries for retrieving NOAA oceanic station data.
