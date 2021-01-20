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
gocean tides [station-id]
```

For example:

```bash
$ gocean tides 9410170
  2021-01-19 09:54      2.732
  2021-01-19 10:00      2.765
  2021-01-19 10:06      2.798
  2021-01-19 10:12      2.831
  2021-01-19 10:18      2.864
  2021-01-19 10:24      2.898
  2021-01-19 10:30      2.931
  2021-01-19 10:36      2.964
  2021-01-19 10:42      2.998
  2021-01-19 10:48      3.030
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
