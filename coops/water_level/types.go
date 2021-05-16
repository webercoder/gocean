package water_level

import "github.com/webercoder/gocean/coops/coops_client"

type WaterLevelAPI struct {
	App    string
	Client *coops_client.Client
}

type WaterLevels struct {
	WaterLevels []WaterLevel `json:"data"`
}

type WaterLevel struct {
	Time    string `json:"t"`
	Value   string `json:"v"`
	Quality string `json:"q"`
}
