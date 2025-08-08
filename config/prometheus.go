package config

import "github.com/prometheus/client_golang/prometheus"

var GaugeMetrics = struct {
	Temperature, Humidity *prometheus.GaugeVec
}{
	Temperature: prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "temperature_celsius",
			Help: "Current temperature in degrees Celsius",
		},
		[]string{"device_id"},
	),
	Humidity: prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "humidity_percent",
			Help: "Relative humidity in percent",
		},
		[]string{"device_id"},
	),
}
