# gocean

gocean is a set of NOAA API wrappers and tools written in Go. It is currently under initial development
and not ready for external consumption.

## Installation

Retrieve the src and install using the following commands.

```bash
go get github.com/webercoder/gocean/...
go install github.com/webercoder/gocean
```

## Usage

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
  2021-05-14 22:24\t5.174
  2021-05-14 22:30\t5.187
  2021-05-14 22:36\t5.196
  2021-05-14 22:42\t5.201
  2021-05-14 22:48\t5.201
...
```

## Development Plan

Directories under /noaa represent actual NOAA APIs and will be
consumable by external programs. Currently these include partial implementations for the
NOAA CO-OPS Tides and Currents API and a pseudo-API wrapper for the listing of NOAA
oceanic stations. Other APIs will be added in the future.

The top-level program allows users to query the APIs.

Planned features for the top-level executable:

* Complete query capability for all APIs in the /noaa directory.
* Save preferred station(s) to ~/.gocean/ or a similar location.
* Optimize performance of nearest station calculation.
* Allow querying by any location or postcode directly when using the tides and currents API.

Planned features for the APIs:

* Implement the complete NOAA Tides and Currents API.
* Allow the stations listing to be cached locally.
* Add additional NOAA API wrappers.
