package http_status

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type StatusOK struct {
	Status   uint16      `json:"status" example:"200"`
	Message  string      `json:"message" example:"Data returned correctly"`
	Response interface{} `json:"response,omitempty"`
} //@name StatusOK

type StatusCreated struct {
	Status   uint16      `json:"status" example:"201"`
	Message  string      `json:"message" example:"Entity created"`
	Response interface{} `json:"response,omitempty"`
} //@name StatusCreated

type StatusNoContent struct {
	Status   uint16      `json:"status" example:"204"`
	Message  string      `json:"message" example:"No content"`
	Response interface{} `json:"response,omitempty"`
} //@name StatusNoContent

func (h *HttpStatus) NoContent(ctx *gin.Context) {
	var response StatusNoContent
	empty := EmptyInterface{}

	response.Status = http.StatusNoContent
	response.Message = "No Content"
	response.Response = empty

	ctx.AbortWithStatusJSON(http.StatusNoContent, response)
}
