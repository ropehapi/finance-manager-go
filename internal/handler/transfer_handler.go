package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/ropehapi/finance-manager-go/internal/model"
	"github.com/ropehapi/finance-manager-go/internal/service"
)

type TransferHandler struct {
	svc service.TransferService
}

func NewTransferHandler(svc service.TransferService) *TransferHandler {
	return &TransferHandler{svc}
}

func (h *TransferHandler) RegisterRoutes(r chi.Router) {
	r.Route("/transfers", func(r chi.Router) {
		r.Post("/", h.Create)
		r.Get("/", h.GetAll)
		r.Get("/{id}", h.GetByID)
		r.Delete("/{id}", h.Delete)
	})
}

func (h *TransferHandler) Create(w http.ResponseWriter, r *http.Request) {
	var input model.Transfer
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	transfer, err := h.svc.Create(r.Context(), input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(transfer)
}

func (h *TransferHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	transfers, err := h.svc.GetAll(r.Context())
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(transfers)
}

func (h *TransferHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	_, err := uuid.Parse(idParam)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	transfer, err := h.svc.GetByID(r.Context(), idParam)
	if err != nil {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(transfer)
}

func (h *TransferHandler) Delete(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	_, err := uuid.Parse(idParam)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	if err := h.svc.Delete(r.Context(), idParam); err != nil {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
