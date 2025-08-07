package domain

type Data struct {
	ID          int     `json:"id,omitempty"`
	Temperature float64 `json:"temperature,omitempty"`
	Humidity    float64 `json:"humidity,omitempty"`
	CreatedAt   string  `json:"created_at,omitempty"`
}
