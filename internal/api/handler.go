package api

import (
	"TestTask/pkg/logging"
	"TestTask/pkg/postgresql"
	"net/http"
)

type Handler struct {
	client postgresql.Client
	logger *logging.Logger
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	// TODO
}

func (h *Handler) List(w http.ResponseWriter, r *http.Request) {
	// TODO
}

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	// TODO
}

func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	// TODO
}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	// TODO
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
