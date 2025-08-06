package handler

import (
	"net/http"
	"strconv"

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
	limitStr := c.DefaultQuery("limit", "10")
	offsetStr := c.DefaultQuery("offset", "0")

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit <= 0 {
		limit = 10
	}

	offset, err := strconv.Atoi(offsetStr)
	if err != nil || offset < 0 {
		offset = 0
	}

	name := c.Query("name")
	currencyCode := c.Query("currency_code")
	accountId := c.Query("account_id")

	filter := model.PaymentMethodFilter{
		Name:      name,
		Type:      currencyCode,
		AccountID: accountId,
		Limit:     limit,
		Offset:    offset,
	}

	output, err := h.svc.GetAll(c.Request.Context(), filter)
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
