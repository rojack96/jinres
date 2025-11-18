package jinres

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rojack96/jinres/http_status"
)

type Customs struct{}

type EmptyInterface struct {
	Empty bool `json:"empty,omitempty"`
}

func (c *Customs) UniqueIdValidator(ctx *gin.Context) {
	var response http_status.StatusBadRequest
	empty := &EmptyInterface{}

	response.Status = http.StatusBadRequest
	response.Message = "Unique Id is incorrect"
	response.Response = empty

	ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
}

func (c *Customs) MakeOrUnitIdNotExist(ctx *gin.Context, mk, imei string) {
	var response http_status.StatusNotFound
	empty := &EmptyInterface{}

	response.Status = http.StatusNotFound
	if mk == "" {
		response.Message = "Error encountered in unit search, details: Make of unit not found"
	} else if imei == "" {
		response.Message = "Error encountered in unit search, details: IMEI of unit not found"
	}
	response.Response = empty

	ctx.AbortWithStatusJSON(http.StatusNotFound, response)
}

func (c *Customs) CommandUnrecognized(ctx *gin.Context) {
	var response http_status.StatusBadRequest
	empty := &EmptyInterface{}

	response.Status = http.StatusBadRequest
	response.Message = "Written command was not recognized"
	response.Response = empty

	ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
}

func (c *Customs) TimestampIncorrectLayout(ctx *gin.Context) {
	var response http_status.StatusBadRequest
	empty := &EmptyInterface{}

	response.Status = http.StatusBadRequest
	response.Message = "Layout of timestamp is incorrect"
	response.Response = empty

	ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
}

func (c *Customs) IdNotValid(ctx *gin.Context) {
	var response http_status.StatusBadRequest
	empty := &EmptyInterface{}

	response.Status = http.StatusBadRequest
	response.Message = "Id not valid"
	response.Response = empty

	ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
}

func (c *Customs) CommandTypeNotValid(ctx *gin.Context) {
	var response http_status.StatusBadRequest
	empty := &EmptyInterface{}

	response.Status = http.StatusBadRequest
	response.Message = "Command type not valid"
	response.Response = empty

	ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
}
