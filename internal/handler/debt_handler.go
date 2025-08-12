package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/ropehapi/finance-manager-go/internal/service"
	"net/http"
)

type DebtHandler struct {
	svc service.DebtService
}

func NewDebtHandler(svc service.DebtService) *DebtHandler {
	return &DebtHandler{svc}
}

func (h *DebtHandler) GetAll(c *gin.Context) {
	output, err := h.svc.GetAll(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	c.JSON(http.StatusOK, output)
}

func (h *DebtHandler) Pay(c *gin.Context) {
	id := c.Param("id")
	payerAccountId := c.Param("payer_account_id")

	output, err := h.svc.Pay(c.Request.Context(), id, payerAccountId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"data": output})
}

func (h *DebtHandler) Delete(c *gin.Context) {
	id := c.Param("id")

	if err := h.svc.Delete(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "fk error"})
		return
	}

	c.Status(http.StatusNoContent)
}
