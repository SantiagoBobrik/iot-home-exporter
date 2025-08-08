package domain

import (
	"errors"
	"strings"
)

type Data struct {
	ID          int     `json:"id,omitempty"`
	DeviceID    string  `json:"device_id,omitempty"`
	Temperature float64 `json:"temperature,omitempty"`
	Humidity    float64 `json:"humidity,omitempty"`
	CreatedAt   string  `json:"created_at,omitempty"`
}

func (d *Data) Validate() error {
	// Validar que DeviceID no esté vacío
	if strings.TrimSpace(d.DeviceID) == "" {
		return errors.New("device_id is required")
	}

	// Validar rango de temperatura (-50 a 100 °C, por ejemplo)
	if d.Temperature < -50 || d.Temperature > 100 {
		return errors.New("temperature must be between -50 and 100 Celsius")
	}

	// Validar rango de humedad (0 a 100 %)
	if d.Humidity < 0 || d.Humidity > 100 {
		return errors.New("humidity must be between 0 and 100 percent")
	}

	return nil
}
