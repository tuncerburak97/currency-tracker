package currency

import "time"

type RateResponse struct {
	Name  string `json:"name"`
	Rates []Rate `json:"rates"`
}

type Rate struct {
	Price     float64   `json:"price"`
	Timestamp time.Time `json:"timestamp"`
}
