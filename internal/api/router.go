package api

import "net/http"

func RegisterRoutes(mux *http.ServeMux, h *Handler) {
	mux.HandleFunc("POST /subscriptions", h.Create)
	mux.HandleFunc("GET /subscriptions", h.List)
	mux.HandleFunc("GET /subscriptions/{id}", h.Get)
	mux.HandleFunc("PUT /subscriptions/{id}", h.Update)
	mux.HandleFunc("DELETE /subscriptions/{id}", h.Delete)
}
