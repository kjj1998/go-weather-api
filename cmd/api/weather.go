package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) getWeatherHandler(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())

	location := params.ByName("location")

	fmt.Fprintf(w, "show the weather details of location: %s\n", location)
}
