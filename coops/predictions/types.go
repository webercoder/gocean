package predictions

import (
	"github.com/webercoder/gocean/coops/coops_client"
)

// Predictions contain tide predictions from a station
type predictions struct {
	Predictions []Prediction `json:"predictions"`
}

// Prediction contains a single tide prediction for a specific time.
type Prediction struct {
	Time  string `json:"t"`
	Value string `json:"v"`
}

type PredictionsApi struct {
	App    string
	Client *coops_client.Client
}
