# gocean

[![<ORG_NAME>](https://circleci.com/gh/webercoder/gocean.svg?style=svg)](https://circleci.com/gh/webercoder/gocean)
[![Go Report Card](https://goreportcard.com/badge/github.com/webercoder/gocean)](https://goreportcard.com/report/github.com/webercoder/gocean)
[![Maintainability](https://api.codeclimate.com/v1/badges/f9d7f2157e1538a06b13/maintainability)](https://codeclimate.com/github/webercoder/gocean/maintainability)

gocean is a set of NOAA API wrappers and tools written in Go. It is currently under active, initial development.

## Table of Contents

* [Installation](#installation)
* [Command-Line Usage](#command-line-usage)
  * [CO-OPS Stations API](#co-ops-stations-api)
    * [Air Gap](#air-gap)
    * [Air Pressure](#air-pressure)
    * [Air Temperatures](#air-temperatures)
    * [Conductivity](#conductivity)
    * [Tide Predictions](#tide-predictions)
    * [Visibility](#visibility)
    * [Water Levels](#water-levels)
    * [Water Temperatures](#water-temperatures)
    * [Wind](#wind)
  * [Stations Info API](#stations-info-api)
    * [Get the Nearest Station](#get-the-nearest-station)
* [Code Structure](#code-structure)

## Installation

To use the library:

```txt
go get -u github.com/webercoder/gocean
```

To install the binary (assuming `GOPATH/bin` is in your `PATH`):

```txt
go install github.com/webercoder/gocean
```

## Command-Line Usage

### CO-OPS Stations API

The operations below use the following command-line parameter structure ([in Go flag syntax](https://golang.org/pkg/flag/#hdr-Command_line_flag_syntax)):

```txt
-begin-date string
    The begin date for the data set.
-count int
    The number of results to display. Only works with the pretty format. (default -1)
-datum string
    The datum to query. Possible values: [CRD IGLD LWD MHHW MHW MTL MSL MLW MLLW NAVD STND] (default "MLLW")
-end-date string
    The end date for the data set.
-format string
    The output format of the results. Possible values: [json xml csv pretty] (default "pretty")
-hours int
    The offset from the start time. (default 24)
-station string
    The station to query.
-time-zone-format string
    The time zone format. Possible values: [gmt lst lst_ldt] (default "lst_ldt")
-units string
    Either english or metric. Possible values: [english metric] (default "english")
```

#### Air Gap

```txt
gocean coops air_gap
```

Example:

```txt
$ gocean coops air_gap -station 8545556 -count 5
Air gap readings for station 8545556:

2021-05-31 06:48: 130.945ft (Sigma: 0.039, Flags: 0,0,0,0)
2021-05-31 06:54: 130.938ft (Sigma: 0.030, Flags: 0,0,0,0)
2021-05-31 07:00: 130.997ft (Sigma: 0.020, Flags: 0,0,0,0)
2021-05-31 07:06: 131.037ft (Sigma: 0.023, Flags: 0,0,0,0)
2021-05-31 07:12: 131.106ft (Sigma: 0.030, Flags: 0,0,0,0)
```

#### Air Pressure

```txt
gocean coops air_pressure
```

Example:

```txt
$ gocean coops air_pressure -station 9410230 -count 5
Air pressure readings for station 9410230:

2021-05-31 06:48: 1014.6mbar (Flags: 0,0,0)
2021-05-31 06:54: 1014.6mbar (Flags: 0,0,0)
2021-05-31 07:00: 1014.7mbar (Flags: 0,0,0)
2021-05-31 07:06: 1014.7mbar (Flags: 0,0,0)
2021-05-31 07:12: 1014.7mbar (Flags: 0,0,0)
```

#### Air Temperatures

```txt
gocean coops air_temperature
```

Example:

```txt
$ gocean coops air_temperature --station 9410230 --count 5
Air temperature readings for station 9410230:

2021-05-31 06:48: 60.3°f
2021-05-31 06:54: 60.1°f
2021-05-31 07:00: 60.3°f
2021-05-31 07:06: 60.1°f
2021-05-31 07:12: 60.1°f
```

#### Conductivity

```txt
gocean coops conductivity
```

Example:

```txt
$ gocean coops conductivity -station 8447386 -count 5
Conductivity readings for station 8447386:

2021-05-31 06:48: 35.81mS/cm (Flags: 0,0,0)
2021-05-31 06:54: 36.23mS/cm (Flags: 0,0,0)
2021-05-31 07:00: 36.44mS/cm (Flags: 0,0,0)
2021-05-31 07:06: 36.51mS/cm (Flags: 0,0,0)
2021-05-31 07:12: 35.84mS/cm (Flags: 0,0,0)
```

#### Visibility

```txt
gocean coops visiblity
```

Example:

```txt
$ gocean coops visibility -station 8447412 -count 5
Visibility readings for station 8447412:

2021-05-31 06:48: 1.86nmi (Flags: 0,0,0)
2021-05-31 06:54: 1.86nmi (Flags: 0,0,0)
2021-05-31 07:00: 1.90nmi (Flags: 0,0,0)
2021-05-31 07:06: 3.34nmi (Flags: 0,0,0)
2021-05-31 07:12: 3.38nmi (Flags: 0,0,0)
```

#### Tide Predictions

```txt
gocean coops predictions
```

Example:

```txt
$ gocean coops predictions -station 9414523 -count 5
Tide predictions for station 9414523:

2021-06-01 06:42: 6.551ft
2021-06-01 06:48: 6.472ft
2021-06-01 06:54: 6.387ft
2021-06-01 07:00: 6.296ft
2021-06-01 07:06: 6.200ft
```

#### Water Levels

```txt
gocean coops water_level
```

Example:

```txt
$ gocean coops water_level -station 9410230 -count 5
Tide water levels for station 9410230:

2021-05-31 06:48: 0.362ft (Quality: Preliminary, Flags: 1,0,0,0)
2021-05-31 06:54: 0.289ft (Quality: Preliminary, Flags: 1,0,0,0)
2021-05-31 07:00: 0.234ft (Quality: Preliminary, Flags: 1,0,0,0)
2021-05-31 07:06: 0.168ft (Quality: Preliminary, Flags: 1,0,0,0)
2021-05-31 07:12: 0.076ft (Quality: Preliminary, Flags: 1,0,0,0)
```

#### Water Temperatures

```txt
gocean coops water_temperature
```

Example:

```txt
$ gocean coops water_temperature --station 9410230 --count 5
Water temperature readings for station 9410230:

2021-05-31 06:48: 66.7°f
2021-05-31 06:54: 66.7°f
2021-05-31 07:00: 66.7°f
2021-05-31 07:06: 66.7°f
2021-05-31 07:12: 66.7°f
```

#### Wind

```txt
gocean coops wind
```

Example:

```txt
$ gocean coops wind -station 9410230 -count 5
Wind readings for station 9410230:

2021-05-31 06:48: 3.89kn from the SW (218.00) with gusts of 5.25kn
2021-05-31 06:54: 2.92kn from the SSW (206.00) with gusts of 5.25kn
2021-05-31 07:00: 2.14kn from the SSW (196.00) with gusts of 4.47kn
2021-05-31 07:06: 5.05kn from the S (188.00) with gusts of 6.61kn
2021-05-31 07:12: 4.67kn from the SSW (198.00) with gusts of 7.00kn
```

### Stations Info API

This isn't an official NOAA API, but it uses the stations list to provide additional information.

#### Get the Nearest Station

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

## Code Structure

Directory | Description
--- | ---
`/command` | Command-line options for the main program.
`/src/coops` | Libraries for connecting to the [NOAA CO-OPS API](https://api.tidesandcurrents.noaa.gov/api/prod/).
`/src/stations` | Libraries for retrieving NOAA oceanic station data.
