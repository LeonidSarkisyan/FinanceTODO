package handlers

import (
	"FinanceTODO/internal/handlers/responses"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"net/http"
)

func (h *Handler) GetUserById(c *gin.Context) {
	userID := getUserID(c)

	user, err := h.services.User.GetById(userID)
	if err != nil {
		log.Error().Err(err).Int("user id", userID).Msg("ошибка при получении пользователя")
		responses.NewErrorResponse(
			c, http.StatusInternalServerError, err, err.Error(),
		)
		return
	}

	c.JSON(http.StatusOK, user)
}
