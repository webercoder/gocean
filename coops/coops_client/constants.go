package coops_client

const (
	ResponseFormatJSON ResponseFormat = iota
	ResponseFormatXML
	ResponseFormatCSV
)

func (s ResponseFormat) String() string {
	return [...]string{"json", "xml", "csv"}[s]
}

const (
	TimeZoneFormatGMT TimeZoneFormat = iota
	TimeZoneFormatLST
	TimeZoneFormatLSTLDT
)

func (s TimeZoneFormat) String() string {
	return [...]string{"gmt", "lst", "lst_ldt"}[s]
}

const (
	UnitsEnglish Units = iota
	UnitsMetric
)

func (s Units) String() string {
	return [...]string{"english", "metric"}[s]
}

const (
	DatumCRD  Datum = iota // Columbia River Datum
	DatumIGLD              // International Great Lakes Datum
	DatumLWD               // Great Lakes Low Water Datum (Chart Datum)
	DatumMHHW              // Mean Higher High Water
	DatumMHW               // Mean High Water
	DatumMTL               // Mean Tide Level
	DatumMSL               // Mean Sea Level
	DatumMLW               // Mean Low Water
	DatumMLLW              // Mean Lower Low Water
	DatumNAVD              // North American Vertical Datum
	DatumSTND              // Station Datum
)

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
	ProductAirGap              Product = iota // Air Gap (distance between a bridge and the water's surface) at the station.
	ProductAirPressure                        // Barometric pressure as measured at the station.
	ProductAirTemperature                     // Air temperature as measured at the station.
	ProductConductivity                       // The water's conductivity as measured at the station.
	ProductCurrents                           // Currents data for currents stations.
	ProductCurrentsPredictions                // Currents predictions data for currents predictions stations.
	ProductDailyMean                          // Verified daily mean water level data for the station.
	ProductDatums                             // datums data for the stations.
	ProductHighLow                            // Verified high/low water level data for the station.
	ProductHourlyHeight                       // Verified hourly height water level data for the station.
	ProductHumidity                           // Relative humidity as measured at the station.
	ProductMonthlyMean                        // Verified monthly mean water level data for the station.
	ProductOneMinuteWaterLevel                // One minute water level data for the station.
	ProductPredictions                        // 6 minute predictions water level data for the station.*
	ProductSalinity                           // Salinity and specific gravity data for the station.
	ProductVisibility                         // Visibility from the station's visibility sensor. A measure of atmospheric clarity.
	ProductWaterLevel                         // Preliminary or verified water levels, depending on availability.
	ProductWaterTemperature                   // Water temperature as measured at the station.
	ProductWind                               // Wind speed, direction, and gusts as measured at the station.
)

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
		"one_minute_water_level",
		"predictions",
		"salinity",
		"visibility",
		"water_level",
		"water_temperature",
		"wind",
	}[s]
}
