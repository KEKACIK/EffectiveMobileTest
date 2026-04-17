package api

import (
	"TestTask/internal/subscriptions"
	"TestTask/internal/validation"
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

func NewHandler(
	client postgresql.Client,
	logger *logging.Logger,
) *Handler {

	return &Handler{
		client: client,
		logger: logger,
	}
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var req SubscriptionCreateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		SendErrorResponse(w, http.StatusBadRequest, err)

		return
	}

	dtoList, err := GetSubscriptionCreateDTO(&req)

	if err != nil {
		SendErrorResponse(w, http.StatusBadRequest, err)

		return
	}

	subRepo := subscriptions.NewRepository(h.client, h.logger)

	subList := make([]*subscriptions.Subscription, 0)
	for _, dto := range dtoList {
		sub, err := subRepo.Create(context.TODO(), dto)
		if err != nil {
			SendErrorResponse(w, http.StatusBadRequest, err)

			return
		}

		subList = append(subList, sub)
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(subList)
}

func (h *Handler) List(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	query := r.URL.Query()
	req, err := GetSubscriptionListDTO(query.Get("page"), query.Get("limit"))
	if err != nil {
		SendErrorResponse(w, http.StatusBadRequest, err)

		return
	}

	subRepo := subscriptions.NewRepository(h.client, h.logger)
	subs, err := subRepo.GetList(context.TODO(), false)
	if err != nil {
		SendErrorResponse(w, http.StatusBadRequest, err)

		return
	}

	totalItems := len(subs)
	items := make([]any, 0)
	for _, sub := range subs {
		items = append(items, sub)
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(GetPagination(items, totalItems, req.Page, req.Limit))
}

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, err := validation.SubscriptionIdValidate(r.PathValue("id"))
	if err != nil {
		SendErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	subRepo := subscriptions.NewRepository(h.client, h.logger)
	sub, err := subRepo.Get(context.TODO(), id, false)
	if err != nil {
		SendErrorNotFoundResponse(w)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(sub)
}

func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Получаем/проверяем ID
	id, err := validation.SubscriptionIdValidate(r.PathValue("id"))
	if err != nil {
		SendErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	// Получаем данные
	var req SubscriptionUpdateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		SendErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	// Проверяем данные
	dto, err := GetSubscriptionUpdateDTO(id, &req)
	if err != nil {
		if err == SubscriptionUpdateEmptyErr {
			SendErrorResponse(w, http.StatusNoContent, err)
			return
		}

		SendErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	subRepo := subscriptions.NewRepository(h.client, h.logger)
	sub, err := subRepo.Update(context.TODO(), dto)
	if err != nil {
		SendErrorNotFoundResponse(w)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(sub)
}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, err := validation.SubscriptionIdValidate(r.PathValue("id"))
	if err != nil {
		SendErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	subRepo := subscriptions.NewRepository(h.client, h.logger)
	err = subRepo.Delete(context.TODO(), id)
	if err != nil {
		SendErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{})
}

func (h *Handler) Stats(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	query := r.URL.Query()

	dto, err := GetSubscriptionStatsDTO(query.Get("name"), query.Get("user_id"), query.Get("start_date"), query.Get("stop_date"))
	if err != nil {
		SendErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	subRepo := subscriptions.NewRepository(h.client, h.logger)
	sum, err := subRepo.Sum(context.TODO(), dto, false)
	if err != nil {
		SendErrorNotFoundResponse(w)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]int{"summary": sum})
}
