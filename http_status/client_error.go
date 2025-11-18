package http_status

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type StatusBadRequest struct {
	Status   uint16      `json:"status" example:"400"`
	Message  string      `json:"message" example:"Bad request"`
	Response interface{} `json:"response,omitempty"`
} //@name StatusBadRequest

type StatusUnauthorized struct {
	Status   uint16      `json:"status" example:"401"`
	Message  string      `json:"message" example:"Unauthorized"`
	Response interface{} `json:"response,omitempty"`
} //@name StatusNotFound

type StatusNotFound struct {
	Status   uint16      `json:"status" example:"404"`
	Message  string      `json:"message" example:"Not found"`
	Response interface{} `json:"response,omitempty"`
} //@name StatusNotFound

func (h *HttpStatus) BadRequest(ctx *gin.Context) {
	var response StatusBadRequest
	empty := EmptyInterface{}

	response.Status = http.StatusBadRequest
	response.Message = "Invalid request message"
	response.Response = empty

	ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
}

func (h *HttpStatus) Unauthorized(ctx *gin.Context) {
	var response StatusUnauthorized
	empty := EmptyInterface{}

	response.Status = http.StatusUnauthorized
	response.Message = "Invalid credential"
	response.Response = empty

	ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
}
