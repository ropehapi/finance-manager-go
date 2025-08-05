package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ropehapi/finance-manager-go/internal/model"
	"github.com/ropehapi/finance-manager-go/internal/service"
)

type PaymentMethodHandler struct {
	svc service.PaymentMethodService
}

func NewPaymentMethodHandler(svc service.PaymentMethodService) *PaymentMethodHandler {
	return &PaymentMethodHandler{svc}
}

func (h *PaymentMethodHandler) Create(c *gin.Context) {
	var input model.CreatePaymentMethodInputDTO

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	output, err := h.svc.Create(c.Request.Context(), input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, output)
}

func (h *PaymentMethodHandler) GetAll(c *gin.Context) {
	output, err := h.svc.GetAll(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	c.JSON(http.StatusOK, output)
}

func (h *PaymentMethodHandler) GetByID(c *gin.Context) {
	id := c.Param("id")

	output, err := h.svc.GetByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}

	c.JSON(http.StatusOK, output)
}

func (h *PaymentMethodHandler) Update(c *gin.Context) {
	id := c.Param("id")
	var input model.UpdatePaymentMethodInputDTO

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	output, err := h.svc.Update(c.Request.Context(), id, input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, output)
}

func (h *PaymentMethodHandler) Delete(c *gin.Context) {
	id := c.Param("id")

	if err := h.svc.Delete(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "fk error"})
		return
	}

	c.Status(http.StatusNoContent)
}
