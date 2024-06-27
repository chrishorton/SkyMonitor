# Nearby Flight Utility

This Go application provides functionality to retrieve information about nearby flights, airport arrivals, and departures using the OpenSky Network API.

## Features

- Fetch nearby flights based on latitude and longitude
- Retrieve arrival information for a specific airport
- Retrieve departure information for a specific airport

## Prerequisites

- Go 1.21.6 or higher
- OpenSky Network API credentials

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/yourusername/nearby_flight_utility.git
   cd nearby_flight_utility
   ```

2. Install dependencies:

   ```bash
   go mod tidy
   ```

## Usage

The application supports three modes of operation:

1. Nearby flights
2. Airport arrivals
3. Airport departures

### Setting up API credentials

Before running the application, you need to set up your OpenSky Network API credentials. The application will prompt you to enter these if they're not already set as environment variables.

### Running the application

Flags:

- `-mode`: Operation mode (nearby, arrivals, or departures)
- `-lat`: Latitude (required for nearby mode)
- `-lon`: Longitude (required for nearby mode)
- `-icao`: ICAO code of the airport (required for arrivals and departures modes)

Examples:

1. Nearby flights:
```go run main.go [flags]```
2. Airport arrivals:
```go run main.go -mode arrivals -icao EHAM```
3. Airport departures:
```go run main.go -mode departures -icao EHAM```

## Code Structure

- `main.go`: Contains the main application logic and CLI interface
- `nearme/nearme.go`: Implements the API request functions for different modes

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

[MIT License](LICENSE)
