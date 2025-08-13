package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"

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

func (h *TransferHandler) Cashin(c *gin.Context) {
	var input model.CreateCashinTransferInputDTO

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	output, err := h.svc.Cashin(c.Request.Context(), input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, output)
}

func (h *TransferHandler) Cashout(c *gin.Context) {
	var input model.CreateCashoutTransferInputDTO

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	output, err := h.svc.Cashout(c.Request.Context(), input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, output)
}

func (h *TransferHandler) GetAll(c *gin.Context) {
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

	transferType := c.Query("type")
	category := c.Query("category")

	filter := model.TransferFilter{
		Type:     transferType,
		Category: category,
		Limit:    limit,
		Offset:   offset,
	}

	output, err := h.svc.GetAll(c.Request.Context(), filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}
	c.JSON(http.StatusOK, output)
}

func (h *TransferHandler) GetByID(c *gin.Context) {
	idParam := c.Param("id")

	if _, err := uuid.Parse(idParam); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	output, err := h.svc.GetByID(c.Request.Context(), idParam)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}

	c.JSON(http.StatusOK, output)
}

func (h *TransferHandler) Delete(c *gin.Context) {
	idParam := c.Param("id")

	if _, err := uuid.Parse(idParam); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	if err := h.svc.Delete(c.Request.Context(), idParam); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}

	c.Status(http.StatusNoContent)
}
