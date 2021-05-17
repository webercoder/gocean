package predictions

import (
	"github.com/webercoder/gocean/coops/coopsclient"
)

// Predictions contain tide predictions from a station
type Predictions struct {
	Predictions []Prediction `json:"predictions"`
}

// Prediction contains a single tide prediction for a specific time.
type Prediction struct {
	Time  string `json:"t"`
	Value string `json:"v"`
}

// API interacts with the predictions product.
type API struct {
	App    string
	Client *coopsclient.Client
}
