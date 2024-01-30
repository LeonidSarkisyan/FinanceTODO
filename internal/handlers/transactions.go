package handlers

import (
	"FinanceTODO/internal/handlers/responses"
	"FinanceTODO/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) CreateTransaction(c *gin.Context) {
	userID := getUserID(c)

	var input models.TransactionInput

	if err := c.BindJSON(&input); err != nil {
		responses.NewErrorResponse(c, http.StatusBadRequest, err, err.Error())
		return
	}

	balanceID, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		responses.NewErrorResponse(c, http.StatusBadRequest, err, err.Error())
		return
	}

	id, err := h.services.Transaction.Create(input, balanceID, input.SubCategoryID, userID)

	if err != nil {
		responses.NewErrorResponse(c, http.StatusInternalServerError, err, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": id})
}

func (h *Handler) GetTransactions(c *gin.Context) {
	userID := getUserID(c)

	balanceID, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		responses.NewErrorResponse(c, http.StatusBadRequest, err, err.Error())
		return
	}

	transactions, err := h.services.Transaction.GetAll(balanceID, userID)

	if err != nil {
		responses.NewErrorResponse(c, http.StatusInternalServerError, err, err.Error())
		return
	}

	c.JSON(http.StatusOK, transactions)
}

func (h *Handler) GetTransactionById(c *gin.Context) {
	userID := getUserID(c)

	transactionID, err := strconv.Atoi(c.Param("transaction_id"))

	if err != nil {
		responses.NewErrorResponse(c, http.StatusBadRequest, err, err.Error())
		return
	}

	transaction, err := h.services.Transaction.GetByID(transactionID, userID)

	if err != nil {
		responses.NewErrorResponse(c, http.StatusInternalServerError, err, err.Error())
		return
	}

	c.JSON(http.StatusOK, transaction)
}

func (h *Handler) UpdateTransaction(c *gin.Context) {
	userID := getUserID(c)

	transactionID, err := strconv.Atoi(c.Param("transaction_id"))

	if err != nil {
		responses.NewErrorResponse(c, http.StatusBadRequest, err, err.Error())
		return
	}

	var input models.TransactionUpdate

	if err = c.BindJSON(&input); err != nil {
		responses.NewErrorResponse(c, http.StatusBadRequest, err, err.Error())
		return
	}

	err = h.services.Transaction.Update(input, transactionID, userID)

	if err != nil {
		responses.NewErrorResponse(c, http.StatusInternalServerError, err, err.Error())
		return
	}

	c.JSON(http.StatusNoContent, gin.H{"status": "ok"})
}

func (h *Handler) DeleteTransaction(c *gin.Context) {
	userID := getUserID(c)

	transactionID, err := strconv.Atoi(c.Param("transaction_id"))

	if err != nil {
		responses.NewErrorResponse(c, http.StatusBadRequest, err, err.Error())
		return
	}

	err = h.services.Transaction.Delete(transactionID, userID)

	if err != nil {
		responses.NewErrorResponse(c, http.StatusInternalServerError, err, err.Error())
		return
	}

	c.JSON(http.StatusNoContent, gin.H{"status": "ok"})
}
