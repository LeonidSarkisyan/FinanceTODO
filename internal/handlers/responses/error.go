package responses

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type ErrorResponse struct {
	Detail string `json:"detail"`
	Error  string `json:"error"`
}

func NewErrorResponse(c *gin.Context, statusCode int, err error, message string) {
	log.Error().Err(err).Msg(message)
	switch {
	case err.Error() == "EOF":
		c.AbortWithStatusJSON(statusCode, ErrorResponse{
			Detail: "пустой request body",
			Error:  err.Error(),
		})
	default:
		c.AbortWithStatusJSON(statusCode, ErrorResponse{
			Detail: message,
			Error:  err.Error(),
		})
	}
}
