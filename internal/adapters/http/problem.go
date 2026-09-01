package http

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/bdobrica/ThinkPixelMEM/internal/domain"
)

var errNotReady = errors.New("service is not ready")

type Problem struct {
	Type      string       `json:"type"`
	Title     string       `json:"title"`
	Status    int          `json:"status"`
	Detail    string       `json:"detail,omitempty"`
	Instance  string       `json:"instance,omitempty"`
	RequestID string       `json:"request_id"`
	Errors    []FieldError `json:"errors,omitempty"`
}
type FieldError struct {
	Field  string `json:"field"`
	Reason string `json:"reason"`
}

// WriteProblem emits RFC 7807 without exposing err.Error() or its causes.
func WriteProblem(w http.ResponseWriter, r *http.Request, err error) {
	status, slug, title, detail := problemMapping(err)
	p := Problem{Type: "https://thinkpixel.dev/problems/" + slug, Title: title, Status: status, Detail: detail, Instance: r.URL.Path, RequestID: RequestID(r.Context())}
	w.Header().Set("Content-Type", "application/problem+json")
	w.Header().Set("Cache-Control", "no-store")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(p)
}

func problemMapping(err error) (int, string, string, string) {
	if errors.Is(err, errNotReady) {
		return 503, "not-ready", "Service unavailable", "The service is not ready to accept traffic."
	}
	var tooLarge *http.MaxBytesError
	if errors.As(err, &tooLarge) {
		return 413, "request-too-large", "Request too large", "The request body exceeds the configured limit."
	}
	code, ok := domain.ErrorCodeOf(err)
	if !ok {
		code = domain.CodeInternal
	}
	switch code {
	case domain.CodeInvalid:
		return 400, "invalid-request", "Invalid request", "The request is not valid."
	case domain.CodeNotFound:
		return 404, "not-found", "Not found", "The requested resource was not found."
	case domain.CodeConflict:
		return 409, "conflict", "Conflict", "The request conflicts with current state."
	case domain.CodeUnauthorized:
		return 401, "unauthorized", "Unauthorized", "Authentication is required."
	case domain.CodeForbidden:
		return 403, "forbidden", "Forbidden", "The operation is not permitted."
	case domain.CodeExpired:
		return 410, "expired", "Expired", "The requested resource has expired."
	default:
		return 500, "internal", "Internal server error", "The server could not complete the request."
	}
}
