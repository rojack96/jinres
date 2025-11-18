package http_status

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type StatusInternalServerError struct {
	Status   uint16      `json:"status" example:"500"`
	Message  string      `json:"message" example:"Internal server error"`
	Response interface{} `json:"response,omitempty"`
} //@name StatusInternalServerError

func (h *HttpStatus) InternalServerError(ctx *gin.Context) {
	var response StatusInternalServerError
	empty := EmptyInterface{}

	response.Status = http.StatusInternalServerError
	response.Message = "Internal server error"
	response.Response = empty

	ctx.AbortWithStatusJSON(http.StatusInternalServerError, response)
}
