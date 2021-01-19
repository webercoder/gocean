# gocean

gocean is a command line NOAA tool written in Go.

Only current feature:

```bash
gocean [zip code]  // reports closest NOAA ocean measuring station
```

For example:

```bash
$ gocean 94087
The nearest Station is "Redwood City" (ID: 9414523), which is 23.072995 kms away from 94087.

$ gocean 92101
The nearest Station is "San Diego, San Diego Bay" (ID: 9410170), which is 1.130777 kms away from 92101.
```
