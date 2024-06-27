package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"nearby_flight_utility/nearme"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func setEnvironmentalVariable(envKey string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("%s: ", envKey)
	input, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("Error reading input:", err)
		os.Exit(1)
	}
	input = strings.TrimSuffix(input, "\n")

	if err := os.Setenv(envKey, input); err != nil {
		fmt.Printf("Error setting environment variable %s: %v\n", envKey, err)
		os.Exit(1)
	}

	return input
}

func checkEnvironmentalVariable(envKey string) string {
	value, found := os.LookupEnv(envKey)
	if !found {
		value = setEnvironmentalVariable(envKey)
	}
	return value
}

func main() {
	username := checkEnvironmentalVariable("OPENSKY_USERNAME")
	password := checkEnvironmentalVariable("OPENSKY_PASSWORD")

	// Define command-line flags
	lat := flag.Float64("lat", 0, "Latitude")
	lon := flag.Float64("lon", 0, "Longitude")
	icao := flag.String("icao", "", "ICAO code for airport")
	mode := flag.String("mode", "nearby", "Mode: nearby, arrivals, or departures")

	flag.Parse()

	var resp *http.Response
	var err error

	switch *mode {
	case "nearby":
		if *lat == 0 || *lon == 0 {
			log.Fatal("Latitude and longitude are required for nearby mode")
		}
		resp, err = nearme.CreateOpenSkyRequestAll(strconv.FormatFloat(*lat, 'f', -1, 32), strconv.FormatFloat(*lon, 'f', -1, 32), username, password)
		if err != nil {
			log.Fatalf("Error: %v", err)
		}
	case "arrivals":
		if *icao == "" {
			log.Fatal("ICAO code is required for arrivals mode")
		}
		resp, err = nearme.CreateAirportArrivalsRequest(*icao, username, password)
		if err != nil {
			log.Fatalf("Error: %v", err)
		}
	case "departures":
		if *icao == "" {
			log.Fatal("ICAO code is required for departures mode")
		}
		resp, err = nearme.CreateAirportDepartureRequest(*icao, username, password)
		if err != nil {
			log.Fatalf("Error: %v", err)
		}
	default:
		log.Fatalf("Invalid mode: %s", *mode)
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response body: %v", err)
	}

	var response struct {
		Time   int64           `json:"time"`
		States [][]interface{} `json:"states"`
	}
	err = json.Unmarshal(body, &response)
	if err != nil {
		log.Fatalf("Error parsing JSON: %v", err)
	}

	count := 0
	for _, array := range response.States {
		if len(array) > 1 {
			secondElement := array[1]
			fmt.Printf("%v\n", secondElement)
			count++
		}
	}
	fmt.Printf("Total flights: %d\n", count)

	// Process and print the response
	fmt.Printf("Nearby flights response status: %s\n", resp.Status)
}
