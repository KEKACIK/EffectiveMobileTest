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
	var createReq SubscribeCreateRequest
	if err := json.NewDecoder(r.Body).Decode(&createReq); err != nil {
		debug(h.logger, r.RequestURI, r.Method, r.Host, http.StatusBadRequest)
		ErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	createDTO, err := SubscribeCreateValidation(&createReq)

	if err != nil {
		debug(h.logger, r.RequestURI, r.Method, r.Host, http.StatusBadRequest)
		ErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	subRepo := subscriptions.NewRepository(h.client, h.logger)
	sub, err := subRepo.Create(context.TODO(), createDTO)
	if err != nil {
		debug(h.logger, r.RequestURI, r.Method, r.Host, http.StatusBadRequest)
		ErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	debug(h.logger, r.RequestURI, r.Method, r.Host, http.StatusCreated)
	json.NewEncoder(w).Encode(sub)
}

func (h *Handler) List(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	listReq, err := SubscribeListValidation(query.Get("page"), query.Get("limit"))
	if err != nil {
		debug(h.logger, r.RequestURI, r.Method, r.Host, http.StatusBadRequest)
		ErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	subRepo := subscriptions.NewRepository(h.client, h.logger)
	subs, err := subRepo.GetList(context.TODO(), false)
	if err != nil {
		debug(h.logger, r.RequestURI, r.Method, r.Host, http.StatusBadRequest)
		ErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	totalItems := len(subs)
	items := make([]any, 0)
	for _, sub := range subs {
		items = append(items, sub)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	debug(h.logger, r.RequestURI, r.Method, r.Host, http.StatusOK)
	json.NewEncoder(w).Encode(GetPagination(items, totalItems, listReq.Page, listReq.Limit))
}

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	debug(h.logger, r.RequestURI, r.Method, r.Host, http.StatusCreated)
}

func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	debug(h.logger, r.RequestURI, r.Method, r.Host, http.StatusCreated)
}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	deleteReq, err := SubscribeDeleteValidation(r.PathValue("id"))
	if err != nil {
		debug(h.logger, r.RequestURI, r.Method, r.Host, http.StatusBadRequest)
		ErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	subRepo := subscriptions.NewRepository(h.client, h.logger)
	err = subRepo.Delete(context.TODO(), deleteReq.ID)
	if err != nil {
		debug(h.logger, r.RequestURI, r.Method, r.Host, http.StatusBadRequest)
		ErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	debug(h.logger, r.RequestURI, r.Method, r.Host, http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{})
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
