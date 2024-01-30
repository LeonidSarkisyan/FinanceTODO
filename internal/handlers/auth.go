package handlers

import (
	"FinanceTODO/internal/handlers/responses"
	"FinanceTODO/internal/models"
	"FinanceTODO/pkg/utils"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"net/http"
	"strings"
)

// Register @Summary SignUp
// @Tags auth
// @Description create account
// @ID create-account
// @Accept  json
// @Produce  json
// @Param input body models.UserInput true "account info"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} responses.ErrorResponse
// @Failure 500 {object} responses.ErrorResponse
// @Failure default {object} responses.ErrorResponse
// @Router /auth/login [post]
func (h *Handler) Register(c *gin.Context) {
	var input models.UserInput

	if err := c.BindJSON(&input); err != nil {
		responses.NewErrorResponse(
			c, http.StatusUnprocessableEntity, err, err.Error(),
		)
		return
	}

	id, err := h.services.Auth.Register(input)
	if err != nil {
		responses.NewErrorResponse(
			c, http.StatusInternalServerError, err, err.Error(),
		)
		return
	}

	c.JSON(http.StatusCreated, map[string]interface{}{
		"id": id,
	})

	log.Info().Int("user id", id).Str("username", input.Username).Msg("создан новый пользователь")
}

// Login @Summary SignIn
// @Tags users
// @Description login
// @ID login
// @Accept  json
// @Produce  json
// @Param input body models.UserInput true "credentials"
// @Success 200 {string} string "token"
// @Failure 400,404 {object} responses.ErrorResponse
// @Failure 500 {object} responses.ErrorResponse
// @Failure default {object} responses.ErrorResponse
// @Router /users [post]
func (h *Handler) Login(c *gin.Context) {
	var input models.UserInput

	if err := c.BindJSON(&input); err != nil {
		responses.NewErrorResponse(c, http.StatusBadRequest, err, err.Error())
		return
	}

	token, err := h.services.Auth.Login(input)

	if err != nil {
		responses.NewErrorResponse(c, http.StatusUnprocessableEntity, err, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})

	log.Info().Str("phone", input.Phone).Msg("пользователь вошел в систему")
}

func (h *Handler) CurrentUser(c *gin.Context) {
	header := c.GetHeader("Authorization")
	if header == "" {
		responses.NewErrorResponse(
			c, http.StatusUnauthorized, errors.New("пустой auth заголовок"), "пустой auth заголовок",
		)
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		responses.NewErrorResponse(c, http.StatusUnauthorized, errors.New("некорректный auth заголовок"), "некорректный auth заголовок")
		return
	}

	userId, err := utils.GetUserIDFromToken(headerParts[1])
	if err != nil {
		responses.NewErrorResponse(c, http.StatusUnauthorized, errors.New("некорректный токен"), "некорректный токен")
		return
	}
	c.Set("userId", userId)
}

func getUserID(c *gin.Context) int {
	return int(c.MustGet("userId").(uint))
}
