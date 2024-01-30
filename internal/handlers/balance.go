package handlers

import (
	"FinanceTODO/internal/handlers/responses"
	"FinanceTODO/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) CreateBalance(c *gin.Context) {
	userID := getUserID(c)

	var balance models.BalanceInput

	if err := c.BindJSON(&balance); err != nil {
		responses.NewErrorResponse(c, http.StatusBadRequest, err, err.Error())
		return
	}

	id, err := h.services.Balance.Create(balance, userID)

	if err != nil {
		responses.NewErrorResponse(c, http.StatusInternalServerError, err, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": id})
}

func (h *Handler) GetBalances(c *gin.Context) {
	userID := getUserID(c)

	balances, err := h.services.Balance.GetAll(userID)

	if err != nil {
		responses.NewErrorResponse(c, http.StatusInternalServerError, err, err.Error())
		return
	}

	c.JSON(http.StatusOK, balances)
}

func (h *Handler) GetBalanceById(c *gin.Context) {
	userID := getUserID(c)

	balanceID, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		responses.NewErrorResponse(c, http.StatusBadRequest, err, err.Error())
		return
	}

	balance, err := h.services.Balance.GetByID(balanceID, userID)

	if err != nil {
		responses.NewErrorResponse(c, http.StatusInternalServerError, err, err.Error())
		return
	}

	c.JSON(http.StatusOK, balance)
}

func (h *Handler) UpdateBalance(c *gin.Context) {
	userID := getUserID(c)

	balanceID, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		responses.NewErrorResponse(c, http.StatusBadRequest, err, err.Error())
		return
	}

	var balance models.BalanceUpdate

	if err := c.BindJSON(&balance); err != nil {
		responses.NewErrorResponse(c, http.StatusBadRequest, err, err.Error())
		return
	}

	err = h.services.Balance.Update(balance, balanceID, userID)

	if err != nil {
		responses.NewErrorResponse(c, http.StatusInternalServerError, err, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}

func (h *Handler) DeleteBalance(c *gin.Context) {
	userID := getUserID(c)

	balanceID, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		responses.NewErrorResponse(c, http.StatusBadRequest, err, err.Error())
		return
	}

	err = h.services.Balance.Delete(balanceID, userID)

	if err != nil {
		responses.NewErrorResponse(c, http.StatusInternalServerError, err, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}
