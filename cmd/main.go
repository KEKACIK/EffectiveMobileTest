package main

import (
	"TestTask/internal/api"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	handler := api.NewHandler()

	api.RegisterRoutes(mux, handler)

	http.ListenAndServe(":80", mux)
}
