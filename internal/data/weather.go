package data

import "time"

type Weather struct {
	Temperature    float32   `json:"temperature"`
	TemperatureMax float32   `json:"temperature_max"`
	TemperatureMin float32   `json:"temperature_min"`
	DateTime       time.Time `json:"datetime"`
	FeelsLike      float32   `json:"feels_like"`
	Dew            float32   `json:"dew"`
	Humidity       int       `json:"humidity"`
	Precipitation  float32   `json:"precipitation"`
	Location       string    `json:"location"`
}
