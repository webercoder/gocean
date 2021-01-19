package tides

type TideRetrieverRetrieveOptions struct {
	begin_date string
	end_date   string
	date       string
	rng        string
}

type TideRetriever interface {
	retrieveTides(lat, long float64) TideInfo
}
