package coopsclient

const (
	// ResponseFormatJSON represents the JSON format.
	ResponseFormatJSON ResponseFormat = iota

	// ResponseFormatXML represents the XML format.
	ResponseFormatXML

	// ResponseFormatCSV represents the CSV format.
	ResponseFormatCSV
)

// String returns the response format to a string.
func (s ResponseFormat) String() string {
	return [...]string{"json", "xml", "csv"}[s]
}

const (
	// TimeZoneFormatGMT is for GMT time.
	TimeZoneFormatGMT TimeZoneFormat = iota

	// TimeZoneFormatLST is for the local standard time of the target station.
	TimeZoneFormatLST

	// TimeZoneFormatLSTLDT is for the local standard or daylight time of the target station.
	TimeZoneFormatLSTLDT
)

// String returns the timezone format to a string.
func (s TimeZoneFormat) String() string {
	return [...]string{"gmt", "lst", "lst_ldt"}[s]
}

const (
	// UnitsEnglish represents imperial units.
	UnitsEnglish Units = iota

	// UnitsMetric represents metric units.
	UnitsMetric
)

// String returns the timezone format to a string.
func (s Units) String() string {
	return [...]string{"english", "metric"}[s]
}

// Datums! https://tidesandcurrents.noaa.gov/datum_options.html
const (
	// DatumCRD is for the Columbia River Datum.
	DatumCRD Datum = iota

	// DatumIGLD is for the International Great Lakes Datum.
	DatumIGLD

	// DatumLWD is for the Great Lakes Low Water Datum (Chart Datum).
	DatumLWD

	// DatumMHHW is for the Mean Higher High Water datum.
	DatumMHHW

	// DatumMHW is for the Mean High Water datum.
	DatumMHW

	// DatumMTL is for the Mean Tide Level datum.
	DatumMTL

	// DatumMSL is for the Mean Sea Level datum.
	DatumMSL

	// DatumMLW is for the Mean Low Water datum.
	DatumMLW

	// DatumMLLW is for the Mean Lower Low Water datum.
	DatumMLLW

	// DatumNAVD is for the North American Vertical Datum.
	DatumNAVD

	// DatumSTND is the Station Datum.
	DatumSTND
)

// String returns the datum's string value.
func (s Datum) String() string {
	return [...]string{
		"CRD",
		"IGLD",
		"LWD",
		"MHHW",
		"MHW",
		"MTL",
		"MSL",
		"MLW",
		"MLLW",
		"NAVD",
		"STND",
	}[s]
}

const (
	// ProductAirGap is for the air Gap (distance between a bridge and the water's surface) at the station.
	ProductAirGap Product = iota

	// ProductAirPressure is for the barometric pressure as measured at the station.
	ProductAirPressure

	// ProductAirTemperature is for the air temperature as measured at the station.
	ProductAirTemperature

	// ProductConductivity is for the water's conductivity as measured at the station.
	ProductConductivity

	// ProductCurrents is for the currents data for currents stations.
	ProductCurrents

	// ProductCurrentsPredictions is for the currents predictions data for currents predictions stations.
	ProductCurrentsPredictions

	// ProductDailyMean is for the verified daily mean water level data for the station.
	ProductDailyMean

	// ProductDatums is for the datums data for the stations.
	ProductDatums

	// ProductHighLow is for the verified high/low water level data for the station.
	ProductHighLow

	// ProductHourlyHeight is for the verified hourly height water level data for the station.
	ProductHourlyHeight

	// ProductHumidity is for the relative humidity as measured at the station.
	ProductHumidity

	// ProductMonthlyMean is for the verified monthly mean water level data for the station.
	ProductMonthlyMean

	// ProductOneMinuteWaterLevel is for one minute water level data for the station.
	ProductOneMinuteWaterLevel

	// ProductPredictions is for 6 minute predictions water level data for the station.
	ProductPredictions

	// ProductSalinity is for salinity and specific gravity data for the station.
	ProductSalinity

	// ProductVisibility is for visibility from the station's visibility sensor. A measure of atmospheric clarity.
	ProductVisibility

	// ProductWaterLevel is for preliminary or verified water levels, depending on availability.
	ProductWaterLevel

	// ProductWaterTemperature is for water temperature as measured at the station.
	ProductWaterTemperature

	// ProductWind is for wind speed, direction, and gusts as measured at the station.
	ProductWind
)

// String returns the product's string value.
func (s Product) String() string {
	return [...]string{
		"air_gap",
		"air_pressure",
		"air_temperature",
		"conductivity",
		"currents",
		"currents_predictions",
		"daily_mean",
		"datums",
		"high_low",
		"hourly_height",
		"humidity",
		"monthly_mean",
		"one_minute_waterlevel",
		"predictions",
		"salinity",
		"visibility",
		"water_level",
		"water_temperature",
		"wind",
	}[s]
}
