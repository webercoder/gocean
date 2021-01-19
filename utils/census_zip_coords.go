package utils

import (
	"encoding/csv"
	"fmt"
	"io"
	"strconv"
	"strings"
)

// FindCoordsForZip returns the geocoordinates for a zip code.
func FindCoordsForZip(zip string) (*GeoCoordinates, error) {
	// Uncomment to read the zip codes from a file and stream search them:
	// appPath, err := os.Executable()
	// if err != nil {
	// 	fmt.Printf("Could not find executable path for loading census file: %v\n", err)
	// 	return nil, err
	// }
	// path := filepath.Join(filepath.Dir(appPath), "data/census-2019-zip-coords.txt")
	// file, err := os.Open(path)
	// if err != nil {
	// 	fmt.Printf("Could not open census data file: %v\n", err)
	// 	return nil, err
	// }
	// defer file.Close()
	// parser := csv.NewReader(file)

	parser := csv.NewReader(strings.NewReader(CensusZipCoords))
	parser.Comma = '\t'

	for {
		record, err := parser.Read()
		if err == io.EOF {
			break
		}

		if zip == record[0] {
			lat, err := strconv.ParseFloat(strings.TrimSpace(record[5]), 64)
			if err != nil {
				return nil, fmt.Errorf("Found zip but could not parse latitude from Census data")
			}

			long, err := strconv.ParseFloat(strings.TrimSpace(record[6]), 64)
			if err != nil {
				return nil, fmt.Errorf("Found zip but could not parse longitude from Census data")
			}

			return &GeoCoordinates{
				Lat:  lat,
				Long: long,
			}, nil
		}
	}

	return nil, fmt.Errorf("Zip %s not found", zip)
}
