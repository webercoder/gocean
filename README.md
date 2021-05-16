# gocean

gocean is a set of NOAA API wrappers and tools written in Go. It is currently under initial development
and not ready for external consumption.

## Installation

Retrieve the src and install using the following commands.

```bash
go get github.com/webercoder/gocean/...
go install github.com/webercoder/gocean
```

## Command-Line Usage

### Get the Nearest Station

```bash
gocean stations [postcode]
```

For example:

```bash
$ gocean stations 94087
The nearest Station is "Redwood City" (ID: 9414523), which is 23.072995 kms away from 94087.

$ gocean stations 92101
The nearest Station is "San Diego, San Diego Bay" (ID: 9410170), which is 1.130777 kms away from 92101.
```

### Get Tide Predictions

```bash
gocean tidesandcurrents predictions [station-id]
```

For example:

```bash
$ gocean tidesandcurrents predictions 9410230
Tide predictions for station: 9410230
  2021-05-14 22:06\t5.111
  2021-05-14 22:12\t5.136
  2021-05-14 22:18\t5.157
...
```

## Development Plan

The `coops` directory contains libraries for connecting to the [NOAA CO-OPS API](https://api.tidesandcurrents.noaa.gov/api/prod/).

The `stations` directory contains libraries for retrieving NOAA oceanic station
data.

The top-level program has a few convenience functions for querying the API but will
be expanded in the future to support all NOAA operations.

Planned features:

* Flesh out this list and move it to GitHub issues.
* Complete query capability for all CO-OPS products (See [the NOAA CO-OPS API docs](https://api.tidesandcurrents.noaa.gov/api/prod/)).
* Save preferred station(s) to ~/.gocean/ or a similar location for easier command-line usage.
* Optimize performance of nearest station calculation.
* Allow querying by any location or postcode directly when using the tides and currents API.
* Allow the stations listing to be cached locally.
* Write documentation for the API.
