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
