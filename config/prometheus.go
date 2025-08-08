package config

import "github.com/prometheus/client_golang/prometheus"

var GaugeMetrics = struct {
	Temperature, Humidity *prometheus.GaugeVec
}{
	Temperature: prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "temperature_celsius",
			Help: "Temperatura actual en grados Celsius",
		},
		[]string{"device_id"},
	),
	Humidity: prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "humidity_percent",
			Help: "Humedad relativa en %",
		},
		[]string{"device_id"},
	),
}
