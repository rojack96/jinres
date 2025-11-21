package jinres

import (
	"github.com/gin-gonic/gin"
)

// Jinres is the envelope used for JSON responses when writing to gin.Context.
type Jinres struct {
	Status   int    `json:"status"`
	Message  string `json:"message"`
	Response any    `json:"response,omitempty"`
}

// Options is a small builder used to construct and send a response in a fluent style.
// Usage: jrs := NewJinres(ctx)
//
//	jrs.Continue().Done()
//	jrs.Continue().Message("custom").Response(data).Done()
type Options struct {
	j       *Jinres
	status  int
	message string
	resp    any
	custom  map[string]any
}

// NewJinres returns a Jinres helper bound to the given gin.Context.
func NewJinres() *Jinres {
	return &Jinres{}
}

// Message sets a custom message for the response.
func (o *Options) Message(msg string) *Options {
	o.message = msg
	return o
}

// Response sets the response payload. If nil, the Response field will be omitted
// in the JSON output thanks to `omitempty` on Jinres.Response.
func (o *Options) Response(v any) *Options {
	o.resp = v
	return o
}

func (o *Options) Custom(key string, value any) *Options {
	if o.custom == nil {
		o.custom = make(map[string]any)
	}
	o.custom[key] = value
	return o
}

// Done writes the response to the bound gin.Context and aborts the request.
// It uses the prefilled status/message and any provided response payload.
func (o *Options) Done(ctx *gin.Context) {
	if o == nil || o.j == nil || ctx == nil {
		return
	}

	result := gin.H{
		"status":  o.status,
		"message": o.message,
	}

	if o.resp != nil {
		result["response"] = o.resp
	}

	for k, v := range o.custom {
		result[k] = v
	}

	if o.status >= 400 {
		ctx.AbortWithStatusJSON(o.status, result)
	} else {
		ctx.JSON(o.status, result)
	}
}
