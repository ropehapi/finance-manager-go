package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/ropehapi/finance-manager-go/internal/model"
	"github.com/ropehapi/finance-manager-go/internal/service"
)

type PaymentMethodHandler struct {
	svc service.PaymentMethodService
}

func NewPaymentMethodHandler(svc service.PaymentMethodService) *PaymentMethodHandler {
	return &PaymentMethodHandler{svc}
}

func (h *PaymentMethodHandler) RegisterRoutes(r chi.Router) {
	r.Route("/payment-methods", func(r chi.Router) {
		r.Post("/", h.Create)
		r.Get("/", h.GetAll)
		r.Get("/{id}", h.GetByID)
		r.Put("/{id}", h.Update)
		r.Delete("/{id}", h.Delete)
	})
}

func (h *PaymentMethodHandler) Create(w http.ResponseWriter, r *http.Request) {
	var input model.PaymentMethod
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}
	method, err := h.svc.Create(r.Context(), input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(method)
}

func (h *PaymentMethodHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	methods, _ := h.svc.GetAll(r.Context())
	json.NewEncoder(w).Encode(methods)
}

func (h *PaymentMethodHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	method, err := h.svc.GetByID(r.Context(), id)
	if err != nil {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(method)
}

func (h *PaymentMethodHandler) Update(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var input model.PaymentMethod
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}
	method, err := h.svc.Update(r.Context(), id, input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(method)
}

func (h *PaymentMethodHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if err := h.svc.Delete(r.Context(), id); err != nil {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
