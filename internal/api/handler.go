package api

import (
	"TestTask/internal/subscriptions"
	"TestTask/pkg/logging"
	"TestTask/pkg/postgresql"
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

var (
	SubscriptionNotFound error = errors.New("Subscription not found")
)

type Handler struct {
	client postgresql.Client
	logger *logging.Logger
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var req SubscriptionCreateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		debug(h.logger, r.RequestURI, r.Method, r.Host, http.StatusBadRequest)
		ErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	dto, err := SubscriptionCreateValidation(&req)

	if err != nil {
		debug(h.logger, r.RequestURI, r.Method, r.Host, http.StatusBadRequest)
		ErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	subRepo := subscriptions.NewRepository(h.client, h.logger)
	sub, err := subRepo.Create(context.TODO(), dto)
	if err != nil {
		debug(h.logger, r.RequestURI, r.Method, r.Host, http.StatusBadRequest)
		ErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	debug(h.logger, r.RequestURI, r.Method, r.Host, http.StatusCreated)
	json.NewEncoder(w).Encode(sub)
}

func (h *Handler) List(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	query := r.URL.Query()
	req, err := SubscriptionListValidation(query.Get("page"), query.Get("limit"))
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

	w.WriteHeader(http.StatusCreated)
	debug(h.logger, r.RequestURI, r.Method, r.Host, http.StatusOK)
	json.NewEncoder(w).Encode(GetPagination(items, totalItems, req.Page, req.Limit))
}

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	req, err := SubscriptionGetValidation(r.PathValue("id"))
	if err != nil {
		ErrorResponse(w, http.StatusBadRequest, err)
		debug(h.logger, r.RequestURI, r.Method, r.Host, http.StatusBadRequest)
		return
	}

	subRepo := subscriptions.NewRepository(h.client, h.logger)
	sub, err := subRepo.Get(context.TODO(), req.ID, false)
	if err != nil {
		ErrorNotFoundResponse(w)
		debug(h.logger, r.RequestURI, r.Method, r.Host, http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(sub)
	debug(h.logger, r.RequestURI, r.Method, r.Host, http.StatusOK)
}

func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Получаем/проверяем ID
	id, err := SubscriptionIdValidate(r.PathValue("id"))
	if err != nil {
		debug(h.logger, r.RequestURI, r.Method, r.Host, http.StatusBadRequest)
		ErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	// Получаем данные
	var req SubscriptionUpdateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		debug(h.logger, r.RequestURI, r.Method, r.Host, http.StatusBadRequest)
		ErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	// Проверяем данные
	dto, err := SubscriptionUpdateValidation(id, &req)
	if err != nil {
		ErrorResponse(w, http.StatusBadRequest, err)
		debug(h.logger, r.RequestURI, r.Method, r.Host, http.StatusBadRequest)
		return
	}

	subRepo := subscriptions.NewRepository(h.client, h.logger)
	sub, err := subRepo.Update(context.TODO(), dto)
	if err != nil {
		ErrorNotFoundResponse(w)
		debug(h.logger, r.RequestURI, r.Method, r.Host, http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(sub)
	debug(h.logger, r.RequestURI, r.Method, r.Host, http.StatusOK)
}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, err := SubscriptionIdValidate(r.PathValue("id"))
	if err != nil {
		debug(h.logger, r.RequestURI, r.Method, r.Host, http.StatusBadRequest)
		ErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	subRepo := subscriptions.NewRepository(h.client, h.logger)
	err = subRepo.Delete(context.TODO(), id)
	if err != nil {
		debug(h.logger, r.RequestURI, r.Method, r.Host, http.StatusBadRequest)
		ErrorResponse(w, http.StatusBadRequest, err)
		return
	}

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
