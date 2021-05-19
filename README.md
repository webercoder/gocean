# gocean

[![<ORG_NAME>](https://circleci.com/gh/webercoder/gocean.svg?style=svg)](https://circleci.com/gh/webercoder/gocean)
[![Go Report Card](https://goreportcard.com/badge/github.com/webercoder/gocean)](https://goreportcard.com/report/github.com/webercoder/gocean)
[![Maintainability](https://api.codeclimate.com/v1/badges/f9d7f2157e1538a06b13/maintainability)](https://codeclimate.com/github/webercoder/gocean/maintainability)

gocean is a set of NOAA API wrappers and tools written in Go. It is currently under initial development
and not ready for external consumption.

## Installation

Retrieve the src and install using the following commands.

```txt
go get github.com/webercoder/gocean/...
go install github.com/webercoder/gocean
```

## Command-Line Usage

### Get the Nearest Station

```txt
gocean stations [postcode]
```

For example:

```txt
$ gocean stations 94087
The nearest Station is "Redwood City" (ID: 9414523), which is 23.072995 kms away from 94087.

$ gocean stations 92101
The nearest Station is "San Diego, San Diego Bay" (ID: 9410170), which is 1.130777 kms away from 92101.
```

### Get Tide Predictions

```txt
gocean coops predictions [station-id]
```

For example:

```txt
$ gocean coops predictions 9410230
Tide predictions for station: 9410230
  2021-05-14 22:06    5.111
  2021-05-14 22:12    5.136
  2021-05-14 22:18    5.157
...
```

### Get Water Levels

```txt
gocean coops water_level [station-id]
```

For example:

```txt
$ gocean coops water_level 8454000
Tide water levels for station: 8454000
  2021-05-16 16:54    0.635    p
  2021-05-16 17:00    0.661    p
  2021-05-16 17:06    0.625    p
  2021-05-16 17:12    0.523    p
  2021-05-16 17:18    0.444    p
  2021-05-16 17:24    0.415    p
  2021-05-16 17:30    0.421    p
...
```

## Code Structure

The `src/coops` directory contains libraries for connecting to the [NOAA CO-OPS API](https://api.tidesandcurrents.noaa.gov/api/prod/).

The `src/stations` directory contains libraries for retrieving NOAA oceanic station
data.
