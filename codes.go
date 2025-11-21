package jinres

import "net/http"

// 1xx Informational

// Continue returns an Options builder prefilled for HTTP 100 Continue.
// Note: Go methods must be invoked with parentheses, e.g. jrs.Continue().Done().
func (j *Jinres) Continue() *Options {
	return &Options{j: j, status: http.StatusContinue, message: "Continue"}
}

func (j *Jinres) SwitchingProtocols() *Options {
	return &Options{j: j, status: http.StatusSwitchingProtocols, message: "Switching Protocols"}
}

func (j *Jinres) Processing() *Options {
	return &Options{j: j, status: http.StatusProcessing, message: "Processing"}
}

func (j *Jinres) EarlyHints() *Options {
	return &Options{j: j, status: http.StatusEarlyHints, message: "Early Hints"}
}

// 2xx Success

func (j *Jinres) OK() *Options { return &Options{j: j, status: http.StatusOK, message: "OK"} }
func (j *Jinres) Created() *Options {
	return &Options{j: j, status: http.StatusCreated, message: "Created"}
}
func (j *Jinres) Accepted() *Options {
	return &Options{j: j, status: http.StatusAccepted, message: "Accepted"}
}
func (j *Jinres) NonAuthoritativeInformation() *Options {
	return &Options{j: j, status: http.StatusNonAuthoritativeInfo, message: "Non-Authoritative Information"}
}
func (j *Jinres) NoContent() *Options {
	return &Options{j: j, status: http.StatusNoContent, message: "No Content"}
}
func (j *Jinres) ResetContent() *Options {
	return &Options{j: j, status: http.StatusResetContent, message: "Reset Content"}
}
func (j *Jinres) PartialContent() *Options {
	return &Options{j: j, status: http.StatusPartialContent, message: "Partial Content"}
}
func (j *Jinres) MultiStatus() *Options {
	return &Options{j: j, status: http.StatusMultiStatus, message: "Multi-Status"}
}
func (j *Jinres) AlreadyReported() *Options {
	return &Options{j: j, status: http.StatusAlreadyReported, message: "Already Reported"}
}
func (j *Jinres) IMUsed() *Options {
	return &Options{j: j, status: http.StatusIMUsed, message: "IM Used"}
}

// 3xx Redirection

func (j *Jinres) MultipleChoices() *Options {
	return &Options{j: j, status: http.StatusMultipleChoices, message: "Multiple Choices"}
}
func (j *Jinres) MovedPermanently() *Options {
	return &Options{j: j, status: http.StatusMovedPermanently, message: "Moved Permanently"}
}
func (j *Jinres) Found() *Options { return &Options{j: j, status: http.StatusFound, message: "Found"} }
func (j *Jinres) SeeOther() *Options {
	return &Options{j: j, status: http.StatusSeeOther, message: "See Other"}
}
func (j *Jinres) NotModified() *Options {
	return &Options{j: j, status: http.StatusNotModified, message: "Not Modified"}
}
func (j *Jinres) UseProxy() *Options {
	return &Options{j: j, status: http.StatusUseProxy, message: "Use Proxy"}
}
func (j *Jinres) TemporaryRedirect() *Options {
	return &Options{j: j, status: http.StatusTemporaryRedirect, message: "Temporary Redirect"}
}
func (j *Jinres) PermanentRedirect() *Options {
	return &Options{j: j, status: http.StatusPermanentRedirect, message: "Permanent Redirect"}
}

// 4xx Client Error

func (j *Jinres) BadRequest() *Options {
	return &Options{j: j, status: http.StatusBadRequest, message: "Bad Request"}
}
func (j *Jinres) Unauthorized() *Options {
	return &Options{j: j, status: http.StatusUnauthorized, message: "Unauthorized"}
}
func (j *Jinres) PaymentRequired() *Options {
	return &Options{j: j, status: http.StatusPaymentRequired, message: "Payment Required"}
}
func (j *Jinres) Forbidden() *Options {
	return &Options{j: j, status: http.StatusForbidden, message: "Forbidden"}
}
func (j *Jinres) NotFound() *Options {
	return &Options{j: j, status: http.StatusNotFound, message: "Not Found"}
}
func (j *Jinres) MethodNotAllowed() *Options {
	return &Options{j: j, status: http.StatusMethodNotAllowed, message: "Method Not Allowed"}
}
func (j *Jinres) NotAcceptable() *Options {
	return &Options{j: j, status: http.StatusNotAcceptable, message: "Not Acceptable"}
}
func (j *Jinres) ProxyAuthenticationRequired() *Options {
	return &Options{j: j, status: http.StatusProxyAuthRequired, message: "Proxy Authentication Required"}
}
func (j *Jinres) RequestTimeout() *Options {
	return &Options{j: j, status: http.StatusRequestTimeout, message: "Request Timeout"}
}
func (j *Jinres) Conflict() *Options {
	return &Options{j: j, status: http.StatusConflict, message: "Conflict"}
}
func (j *Jinres) Gone() *Options { return &Options{j: j, status: http.StatusGone, message: "Gone"} }
func (j *Jinres) LengthRequired() *Options {
	return &Options{j: j, status: http.StatusLengthRequired, message: "Length Required"}
}
func (j *Jinres) PreconditionFailed() *Options {
	return &Options{j: j, status: http.StatusPreconditionFailed, message: "Precondition Failed"}
}
func (j *Jinres) ContentTooLarge() *Options {
	return &Options{j: j, status: http.StatusRequestEntityTooLarge, message: "Payload Too Large"}
}
func (j *Jinres) URITooLong() *Options {
	return &Options{j: j, status: http.StatusRequestURITooLong, message: "URI Too Long"}
}
func (j *Jinres) UnsupportedMediaType() *Options {
	return &Options{j: j, status: http.StatusUnsupportedMediaType, message: "Unsupported Media Type"}
}
func (j *Jinres) RangeNotSatisfiable() *Options {
	return &Options{j: j, status: http.StatusRequestedRangeNotSatisfiable, message: "Range Not Satisfiable"}
}
func (j *Jinres) ExpectationFailed() *Options {
	return &Options{j: j, status: http.StatusExpectationFailed, message: "Expectation Failed"}
}
func (j *Jinres) ImATeapot() *Options {
	return &Options{j: j, status: http.StatusTeapot, message: "I'm a teapot"}
}
func (j *Jinres) MisdirectedRequest() *Options {
	return &Options{j: j, status: http.StatusMisdirectedRequest, message: "Misdirected Request"}
}
func (j *Jinres) UnprocessableEntity() *Options {
	return &Options{j: j, status: http.StatusUnprocessableEntity, message: "Unprocessable Entity"}
}
func (j *Jinres) Locked() *Options {
	return &Options{j: j, status: http.StatusLocked, message: "Locked"}
}
func (j *Jinres) FailedDependency() *Options {
	return &Options{j: j, status: http.StatusFailedDependency, message: "Failed Dependency"}
}
func (j *Jinres) TooEarly() *Options {
	return &Options{j: j, status: http.StatusTooEarly, message: "Too Early"}
}
func (j *Jinres) UpgradeRequired() *Options {
	return &Options{j: j, status: http.StatusUpgradeRequired, message: "Upgrade Required"}
}
func (j *Jinres) PreconditionRequired() *Options {
	return &Options{j: j, status: http.StatusPreconditionRequired, message: "Precondition Required"}
}
func (j *Jinres) TooManyRequests() *Options {
	return &Options{j: j, status: http.StatusTooManyRequests, message: "Too Many Requests"}
}
func (j *Jinres) RequestHeaderFieldsTooLarge() *Options {
	return &Options{j: j, status: http.StatusRequestHeaderFieldsTooLarge, message: "Request Header Fields Too Large"}
}
func (j *Jinres) UnavailableForLegalReasons() *Options {
	return &Options{j: j, status: http.StatusUnavailableForLegalReasons, message: "Unavailable For Legal Reasons"}
}

// 5xx Server Error

func (j *Jinres) InternalServerError() *Options {
	return &Options{j: j, status: http.StatusInternalServerError, message: "Internal Server Error"}
}
func (j *Jinres) NotImplemented() *Options {
	return &Options{j: j, status: http.StatusNotImplemented, message: "Not Implemented"}
}
func (j *Jinres) BadGateway() *Options {
	return &Options{j: j, status: http.StatusBadGateway, message: "Bad Gateway"}
}
func (j *Jinres) ServiceUnavailable() *Options {
	return &Options{j: j, status: http.StatusServiceUnavailable, message: "Service Unavailable"}
}
func (j *Jinres) GatewayTimeout() *Options {
	return &Options{j: j, status: http.StatusGatewayTimeout, message: "Gateway Timeout"}
}
func (j *Jinres) HTTPVersionNotSupported() *Options {
	return &Options{j: j, status: http.StatusHTTPVersionNotSupported, message: "HTTP Version Not Supported"}
}
func (j *Jinres) VariantAlsoNegotiates() *Options {
	return &Options{j: j, status: http.StatusVariantAlsoNegotiates, message: "Variant Also Negotiates"}
}
func (j *Jinres) InsufficientStorage() *Options {
	return &Options{j: j, status: http.StatusInsufficientStorage, message: "Insufficient Storage"}
}
func (j *Jinres) LoopDetected() *Options {
	return &Options{j: j, status: http.StatusLoopDetected, message: "Loop Detected"}
}
func (j *Jinres) NotExtended() *Options {
	return &Options{j: j, status: http.StatusNotExtended, message: "Not Extended"}
}
func (j *Jinres) NetworkAuthenticationRequired() *Options {
	return &Options{j: j, status: http.StatusNetworkAuthenticationRequired, message: "Network Authentication Required"}
}
