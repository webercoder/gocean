package stations_test

import (
	"net/http"
	"testing"

	"github.com/webercoder/gocean/stations"
	"github.com/webercoder/gocean/testutils"
	"github.com/webercoder/gocean/utils"
)

type FakeStationsClient struct {
	Err error
	XML string
}

const SampleStationsData string = `<?xml version="1.0" encoding="ISO-8859-1" ?>
	<stations xmlns="https://stations-prod.co-ops-aws-east1.net/stations/" 
	xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" 
	xsi:schemaLocation="https://stations-prod.co-ops-aws-east1.net/stations/xml_schemas/stations.xsd"> 
	<station name="Nawiliwili" ID="1611400" >
	<metadata>
	<location>
	<lat>21.9544</lat>
	<long>-159.3561</long>
	<state>HI</state>
	</location>
	<date_established>1954-11-24</date_established>
	</metadata>
	<parameter name="Water Level" sensorID="A1" DCP="1" status="1" />
	<parameter name="Winds" sensorID="C1" DCP="3" status="1" />
	<parameter name="Air Temp" sensorID="D1" DCP="3" status="1" />
	<parameter name="Water Temp" sensorID="E1" DCP="1" status="1" />
	<parameter name="Air Pressure" sensorID="F1" DCP="3" status="1" />
	<parameter name="" sensorID="U1" DCP="1" status="1" />
	</station>
	<station name="Honolulu" ID="1612340" >
	<metadata>
	<location>
	<lat>21.3067</lat>
	<long>-157.867</long>
	<state>HI</state>
	</location>
	<date_established>1905-01-01</date_established>
	</metadata>
	<parameter name="Water Level" sensorID="A1" DCP="1" status="1" />
	<parameter name="Winds" sensorID="C1" DCP="1" status="1" />
	<parameter name="Air Temp" sensorID="D1" DCP="1" status="1" />
	<parameter name="Water Temp" sensorID="E1" DCP="1" status="1" />
	<parameter name="Air Pressure" sensorID="F1" DCP="1" status="1" />
	<parameter name="" sensorID="U1" DCP="1" status="1" />
	</station>
	<station name="San Diego, San Diego Bay" ID="9410170">
	<metadata>
	<location>
	<lat>32.7142</lat>
	<long>-117.1736</long>
	<state>CA</state>
	</location>
	<date_established>1906-01-26</date_established>
	</metadata>
	<parameter name="Water Temp" sensorID="E1" DCP="1" status="1"/>
	<parameter name="Air Pressure" sensorID="F1" DCP="1" status="1"/>
	<parameter name="" sensorID="U1" DCP="1" status="1"/>
	<parameter name="Water Level" sensorID="Y1" DCP="1" status="1"/>
	</station>
	</stations>`

func (fsc *FakeStationsClient) Get(url string) (resp *http.Response, err error) {
	if fsc.Err != nil {
		return nil, fsc.Err
	}

	return &http.Response{
		Body: testutils.NewStringReadCloser(fsc.XML),
	}, nil
}

func TestNOAAStationClient_GetStations(t *testing.T) {
	retriever := &stations.NOAAStationClient{Client: &FakeStationsClient{XML: SampleStationsData}}
	result := retriever.GetStations(true)

	if len(result) != 3 {
		t.Errorf("Found %d stations instead of the expected 3", len(result))
	}

	s := result[1]
	if s.Name != "Honolulu" {
		t.Errorf("Did not expect name of %s", s.Name)
	}
	if s.ID != 1612340 {
		t.Errorf("Did not expect ID of %d", s.ID)
	}
	if s.Metadata.Location.Lat != 21.3067 {
		t.Errorf("Did not expect lat of %f", s.Metadata.Location.Lat)
	}
	if s.Metadata.Location.Long != -157.867 {
		t.Errorf("Did not expect long of %f", s.Metadata.Location.Long)
	}
	if s.Metadata.Location.State != "HI" {
		t.Errorf("Did not expect State of %s", s.Metadata.Location.State)
	}
	if s.Metadata.DateEstablished != "1905-01-01" {
		t.Errorf("Did not expect date established of %s", s.Metadata.DateEstablished)
	}

	c := s.Capabilities
	if len(c) != 6 {
		t.Errorf("Did not expect %d capabilities", len(c))
	} else {
		if c[0].Name != "Water Level" || c[0].SensorID != "A1" || c[0].DCP != 1 || c[0].Status != 1 {
			t.Errorf("Did not expect capability: (%s, %s, %d, %d)", c[0].Name, c[0].SensorID, c[0].DCP, c[0].Status)
		}
		if c[1].Name != "Winds" || c[1].SensorID != "C1" || c[1].DCP != 1 || c[1].Status != 1 {
			t.Errorf("Did not expect capability: (%s, %s, %d, %d)", c[1].Name, c[1].SensorID, c[1].DCP, c[1].Status)
		}
		if c[2].Name != "Air Temp" || c[2].SensorID != "D1" || c[2].DCP != 1 || c[2].Status != 1 {
			t.Errorf("Did not expect capability: (%s, %s, %d, %d)", c[2].Name, c[2].SensorID, c[2].DCP, c[2].Status)
		}
		if c[3].Name != "Water Temp" || c[3].SensorID != "E1" || c[3].DCP != 1 || c[3].Status != 1 {
			t.Errorf("Did not expect capability: (%s, %s, %d, %d)", c[3].Name, c[3].SensorID, c[3].DCP, c[3].Status)
		}
		if c[4].Name != "Air Pressure" || c[4].SensorID != "F1" || c[4].DCP != 1 || c[4].Status != 1 {
			t.Errorf("Did not expect capability: (%s, %s, %d, %d)", c[4].Name, c[4].SensorID, c[4].DCP, c[4].Status)
		}
		if c[5].Name != "" || c[5].SensorID != "U1" || c[5].DCP != 1 || c[5].Status != 1 {
			t.Errorf("Did not expect capability: (%s, %s, %d, %d)", c[5].Name, c[5].SensorID, c[5].DCP, c[5].Status)
		}
	}
}

func TestNOAAStationClient_GetNearestStation(t *testing.T) {
	retriever := &stations.NOAAStationClient{Client: &FakeStationsClient{XML: SampleStationsData}}
	sanDiegoCityHall := utils.GeoCoordinates{Lat: 32.716868, Long: -117.162837}
	expectedStationID := 9410170
	station, _ := retriever.GetNearestStation(sanDiegoCityHall)
	if station.ID != expectedStationID {
		t.Errorf("Expected %d to equal %d", station.ID, expectedStationID)
	}

	pearlHarbor := utils.GeoCoordinates{Lat: 21.339884, Long: -157.970901}
	expectedStationID = 1612340
	station, _ = retriever.GetNearestStation(pearlHarbor)
	if station.ID != expectedStationID {
		t.Errorf("Expected %d to equal %d", station.ID, expectedStationID)
	}
}
