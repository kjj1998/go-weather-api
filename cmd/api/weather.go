package main

import (
	"net/http"
	"time"

	"github.com/kjj1998/go-weather-api/internal/data"
)

func (app *application) getWeatherHandler(w http.ResponseWriter, r *http.Request) {
	location, err := app.readParam(r, "location")
	if err != nil {
		http.NotFound(w, r)
		return
	}

	date, _ := time.Parse("2006-01-02", "2025-05-09")
	weather := data.Weather{
		Temperature:    29.7,
		TemperatureMax: 33.6,
		TemperatureMin: 26.6,
		FeelsLike:      34,
		Dew:            23.8,
		Humidity:       72,
		Precipitation:  5.4,
		DateTime:       date,
		Location:       location,
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"weather": weather}, nil)
	if err != nil {
		app.logger.Error(err.Error())
		http.Error(w, "The server encountered a problem and could not process your request", http.StatusInternalServerError)
	}
}
