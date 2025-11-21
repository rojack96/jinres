package gen

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

// GenerateStructsFile writes a Go source file at outPath that declares a set
// of struct types (one per HTTP status helper) in package `jinres`.
// Each struct follows the shape:
//
//	type BadRequest struct {
//	    Status  uint16 `json:"status" example:"400"`
//	    Message string `json:"message" example:"Bad Request"`
//	}
//
// outPath should be a writable path (e.g. "status_structs_gen.go").
func GenerateStructsFile(outPath string) error {
	entries := []struct {
		Name string
		Code int
		Msg  string
	}{
		// 1xx
		{"Continue", http.StatusContinue, "Continue"},
		{"SwitchingProtocols", http.StatusSwitchingProtocols, "Switching Protocols"},
		{"Processing", http.StatusProcessing, "Processing"},
		{"EarlyHints", http.StatusEarlyHints, "Early Hints"},
		// 2xx
		{"OK", http.StatusOK, "OK"},
		{"Created", http.StatusCreated, "Created"},
		{"Accepted", http.StatusAccepted, "Accepted"},
		{"NonAuthoritativeInformation", http.StatusNonAuthoritativeInfo, "Non-Authoritative Information"},
		{"NoContent", http.StatusNoContent, "No Content"},
		{"ResetContent", http.StatusResetContent, "Reset Content"},
		{"PartialContent", http.StatusPartialContent, "Partial Content"},
		{"MultiStatus", http.StatusMultiStatus, "Multi-Status"},
		{"AlreadyReported", http.StatusAlreadyReported, "Already Reported"},
		{"IMUsed", http.StatusIMUsed, "IM Used"},
		// 3xx
		{"MultipleChoices", http.StatusMultipleChoices, "Multiple Choices"},
		{"MovedPermanently", http.StatusMovedPermanently, "Moved Permanently"},
		{"Found", http.StatusFound, "Found"},
		{"SeeOther", http.StatusSeeOther, "See Other"},
		{"NotModified", http.StatusNotModified, "Not Modified"},
		{"UseProxy", http.StatusUseProxy, "Use Proxy"},
		{"TemporaryRedirect", http.StatusTemporaryRedirect, "Temporary Redirect"},
		{"PermanentRedirect", http.StatusPermanentRedirect, "Permanent Redirect"},
		// 4xx
		{"BadRequest", http.StatusBadRequest, "Bad Request"},
		{"Unauthorized", http.StatusUnauthorized, "Unauthorized"},
		{"PaymentRequired", http.StatusPaymentRequired, "Payment Required"},
		{"Forbidden", http.StatusForbidden, "Forbidden"},
		{"NotFound", http.StatusNotFound, "Not Found"},
		{"MethodNotAllowed", http.StatusMethodNotAllowed, "Method Not Allowed"},
		{"NotAcceptable", http.StatusNotAcceptable, "Not Acceptable"},
		{"ProxyAuthenticationRequired", http.StatusProxyAuthRequired, "Proxy Authentication Required"},
		{"RequestTimeout", http.StatusRequestTimeout, "Request Timeout"},
		{"Conflict", http.StatusConflict, "Conflict"},
		{"Gone", http.StatusGone, "Gone"},
		{"LengthRequired", http.StatusLengthRequired, "Length Required"},
		{"PreconditionFailed", http.StatusPreconditionFailed, "Precondition Failed"},
		{"PayloadTooLarge", http.StatusRequestEntityTooLarge, "Payload Too Large"},
		{"URITooLong", 414, "URI Too Long"},
		{"UnsupportedMediaType", http.StatusUnsupportedMediaType, "Unsupported Media Type"},
		{"RangeNotSatisfiable", http.StatusRequestedRangeNotSatisfiable, "Range Not Satisfiable"},
		{"ExpectationFailed", http.StatusExpectationFailed, "Expectation Failed"},
		{"ImATeapot", http.StatusTeapot, "I'm a teapot"},
		{"MisdirectedRequest", http.StatusMisdirectedRequest, "Misdirected Request"},
		{"UnprocessableEntity", http.StatusUnprocessableEntity, "Unprocessable Entity"},
		{"Locked", http.StatusLocked, "Locked"},
		{"FailedDependency", http.StatusFailedDependency, "Failed Dependency"},
		{"TooEarly", http.StatusTooEarly, "Too Early"},
		{"UpgradeRequired", http.StatusUpgradeRequired, "Upgrade Required"},
		{"PreconditionRequired", http.StatusPreconditionRequired, "Precondition Required"},
		{"TooManyRequests", http.StatusTooManyRequests, "Too Many Requests"},
		{"RequestHeaderFieldsTooLarge", http.StatusRequestHeaderFieldsTooLarge, "Request Header Fields Too Large"},
		{"UnavailableForLegalReasons", http.StatusUnavailableForLegalReasons, "Unavailable For Legal Reasons"},
		// 5xx
		{"InternalServerError", http.StatusInternalServerError, "Internal Server Error"},
		{"NotImplemented", http.StatusNotImplemented, "Not Implemented"},
		{"BadGateway", http.StatusBadGateway, "Bad Gateway"},
		{"ServiceUnavailable", http.StatusServiceUnavailable, "Service Unavailable"},
		{"GatewayTimeout", http.StatusGatewayTimeout, "Gateway Timeout"},
		{"HTTPVersionNotSupported", http.StatusHTTPVersionNotSupported, "HTTP Version Not Supported"},
		{"VariantAlsoNegotiates", http.StatusVariantAlsoNegotiates, "Variant Also Negotiates"},
		{"InsufficientStorage", http.StatusInsufficientStorage, "Insufficient Storage"},
		{"LoopDetected", http.StatusLoopDetected, "Loop Detected"},
		{"NotExtended", http.StatusNotExtended, "Not Extended"},
		{"NetworkAuthenticationRequired", http.StatusNetworkAuthenticationRequired, "Network Authentication Required"},
	}

	var b strings.Builder
	b.WriteString("package jinres\n\n")

	for _, e := range entries {
		// struct comment
		fmt.Fprintf(&b, "// %s represents HTTP %d responses.\n", e.Name, e.Code)
		fmt.Fprintf(&b, "type %s struct {\n", e.Name)
		fmt.Fprintf(&b, "    Status uint16 `json:\"status\" example:\"%d\"`\n", e.Code)
		msgEsc := strings.ReplaceAll(e.Msg, "\"", "\\\"")
		fmt.Fprintf(&b, "    Message string `json:\"message\" example:\"%s\"`\n", msgEsc)
		fmt.Fprintf(&b, "    Response any `json:\"response,omitempty\"`\n")
		fmt.Fprintf(&b, "} ")

		// Swag name comment: struct name + "Response"
		swagName := e.Name + "Response"
		fmt.Fprintf(&b, "// @name %s\n\n", swagName)
	}

	if err := os.WriteFile(outPath, []byte(b.String()), 0644); err != nil {
		return fmt.Errorf("write structs file: %w", err)
	}

	return nil
}
