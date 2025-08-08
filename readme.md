# IoT Home API

This project is a simple RESTful API for storing and retrieving IoT sensor data (temperature and humidity) from various devices. It uses Go, Fiber, SQLite, and Prometheus for metrics.

## Features

- **POST /data**: Receive temperature and humidity data from IoT devices.
- **GET /data**: Retrieve all stored sensor data.
- **GET /metrics**: Expose Prometheus metrics for monitoring.
- **SQLite**: Local database for persistent storage.
- **Prometheus**: Gauge metrics for temperature and humidity per device.


## Prometheus Metrics

The API exposes two Prometheus gauge metrics for monitoring sensor data:

- **temperature_celsius**  
  - **Type:** Gauge  
  - **Labels:** `device_id`  
  - **Description:** Current temperature in degrees Celsius for each device.

- **humidity_percent**  
  - **Type:** Gauge  
  - **Labels:** `device_id`  
  - **Description:** Relative humidity in percent for each device.

You can scrape these metrics at the `/metrics` endpoint. Example output:

```
# HELP temperature_celsius Current temperature in degrees Celsius
# TYPE temperature_celsius gauge
temperature_celsius{device_id="device123"} 23.5

# HELP humidity_percent Relative humidity in percent
# TYPE humidity_percent gauge
humidity_percent{device_id="device123"} 60.2
```

## Technologies

- [Go](https://golang.org/)
- [Fiber](https://github.com/gofiber/fiber)
- [SQLite](https://www.sqlite.org/)
- [Prometheus](https://prometheus.io/)

## Getting Started

### Prerequisites

- Go 1.20+
- SQLite3

### Installation

1. Clone the repository:
    ```sh
    git clone https://github.com/yourusername/iot-home.git
    cd iot-home
    ```

2. Install dependencies:
    ```sh
    go mod tidy
    ```

3. Run the application:
    ```sh
    go run main.go
    ```

The API will start on `http://localhost:3000`.

## API Endpoints

### POST /data

Send sensor data from a device.

**Request Body:**
```json
{
  "device_id": "device123",
  "temperature": 23.5,
  "humidity": 60.2
}
```

**Response:**
- `201 Created` on success
- `400 Bad Request` if JSON is invalid

### GET /data

Retrieve all sensor data.

**Response:**
```json
{
  "data": [
    {
      "id": 1,
      "device_id": "device123",
      "temperature": 23.5,
      "humidity": 60.2,
      "created_at": "2025-08-07 01:28:05"
    }
  ]
}
```

### GET /metrics

Prometheus metrics endpoint for monitoring.

## Project Structure

- `main.go` - Application entry point and route definitions
- `domain/` - Data models and response structures
- `db/` - Database logic
- `config/` - Metrics configuration

##
