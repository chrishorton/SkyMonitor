package nearme

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
)

type OpenSkyRequest struct {
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
}

// radius should be 2.25
func CalculateLatLonBoundingBox(lat, lon float32, radius float32) (float32, float32, float32, float32) {
	latMin := lat - radius
	latMax := lat + radius
	lonMin := lon - radius
	lonMax := lon + radius
	return latMin, latMax, lonMin, lonMax
}

func CreateOpenSkyRequestAll(lat, lon, username, password string) (*http.Response, error) {
	latFloat, _ := strconv.ParseFloat(lat, 64)
	lonFloat, _ := strconv.ParseFloat(lon, 64)

	latMin, latMax, lonMin, lonMax := CalculateLatLonBoundingBox(float32(latFloat), float32(lonFloat), 0.5)

	url := fmt.Sprintf("https://opensky-network.org/api/states/all?lamin=%f&lamax=%f&lomin=%f&lomax=%f", latMin, latMax, lonMin, lonMax)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	req.SetBasicAuth(username, password)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making request: %v", err)
	}

	return resp, nil
}

func CreateAirportArrivalsRequest(icao, username, password string) (*http.Response, error) {
	now := time.Now().Add(-10 * time.Minute).Unix()
	begin := now
	end := now + 3600
	url := fmt.Sprintf("https://opensky-network.org/api/flights/arrival?airport=%s&begin=%d&end=%d", icao, begin, end)
	fmt.Printf("url: %s", url)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	req.SetBasicAuth(username, password)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making request: %v", err)
	}

	return resp, nil
}

func CreateAirportDepartureRequest(icao, username, password string) (*http.Response, error) {
	url := fmt.Sprintf("https://opensky-network.org/api/flights/departure?airport=%s&begin=1517227200&end=1517230800", icao)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	req.SetBasicAuth(username, password)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making request: %v", err)
	}

	return resp, nil
}
