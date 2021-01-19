# gocean

gocean is a command line NOAA tool written in Go.

## Current features

### Get Nearest Station

```bash
gocean station [postcode]  // reports closest NOAA ocean measuring station
```

For example:

```bash
$ gocean station 94087
The nearest Station is "Redwood City" (ID: 9414523), which is 23.072995 kms away from 94087.

$ gocean station 92101
The nearest Station is "San Diego, San Diego Bay" (ID: 9410170), which is 1.130777 kms away from 92101.
```

### Get Tide Predictions

***!WARNING! Not currently returning the correct data.***

```bash
gocean tides [station-id]
```

For example:

```bash
$ gocean tides 9410170
  2021-01-19 09:30      6.396
  2021-01-19 09:36      6.427
  2021-01-19 09:42      6.458
  2021-01-19 09:48      6.489
  2021-01-19 09:54      6.522
  2021-01-19 10:00      6.554
  2021-01-19 10:06      6.587
  2021-01-19 10:12      6.620
...
```
