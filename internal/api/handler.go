package api

import (
	"TestTask/internal/subscriptions"
	"TestTask/pkg/logging"
	"TestTask/pkg/postgresql"
	"context"
	"encoding/json"
	"net/http"
)

type Handler struct {
	client postgresql.Client
	logger *logging.Logger
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	debug(h.logger, r.RequestURI, r.Method, r.Host)

	var createReq CreateRequest
	if err := json.NewDecoder(r.Body).Decode(&createReq); err != nil {
		ErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	createDTO, err := CreateRequestValidation(&createReq)

	if err != nil {
		ErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	subRepo := subscriptions.NewRepository(h.client, h.logger)
	sub, err := subRepo.Create(context.TODO(), createDTO)
	if err != nil {
		ErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]any{
		"status":    "ok",
		"subscribe": sub,
	})
}

func (h *Handler) List(w http.ResponseWriter, r *http.Request) {
	debug(h.logger, r.RequestURI, r.Method, r.Host)
}

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	debug(h.logger, r.RequestURI, r.Method, r.Host)
}

func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	debug(h.logger, r.RequestURI, r.Method, r.Host)
}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	debug(h.logger, r.RequestURI, r.Method, r.Host)

	var deleteReq DeleteRequest
	if err := json.NewDecoder(r.Body).Decode(&deleteReq); err != nil {
		ErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	subRepo := subscriptions.NewRepository(h.client, h.logger)
	err := subRepo.Delete(context.TODO(), deleteReq.ID)
	if err != nil {
		ErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]any{
		"status": "ok",
	})
}

func NewHandler(
	client postgresql.Client,
	logger *logging.Logger,
) *Handler {

	return &Handler{
		client: client,
		logger: logger,
	}
}
