package waterlevel

import "github.com/webercoder/gocean/src/coops/coopsclient"

// API interacts with the water level product.
type API struct {
	App    string
	Client *coopsclient.Client
}

// WaterLevels is a deserialized water level response.
type WaterLevels struct {
	WaterLevels []WaterLevel `json:"data"`
}

// WaterLevel is a singular, deserialized water level.
type WaterLevel struct {
	Time    string `json:"t"`
	Value   string `json:"v"`
	Quality string `json:"q"`
}
