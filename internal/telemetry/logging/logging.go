// Package logging creates structured, redaction-safe service loggers.
//
// Log messages must be stable operational descriptions, never runtime payloads.
// Correlation and outcome data belongs in structured attributes.
package logging

import (
	"errors"
	"io"
	"log/slog"
	"strings"
	"time"
)

const redacted = "[REDACTED]"

// Correlation contains the identifiers permitted by the observability contract.
// Tenant must already be hashed when deployment policy requires hashing.
type Correlation struct {
	Tenant           string
	MemorySpaceID    string
	MemoryID         string
	RevisionID       string
	IngestionEventID string
	ExtractionJobID  string
	ContextPackID    string
	RunID            string
	SessionID        string
	WorkspaceID      string
	RequestID        string
	TraceID          string
}

// Attributes converts non-empty correlation values to canonical log fields.
func (c Correlation) Attributes() []slog.Attr {
	values := []struct{ key, value string }{
		{"tenant", c.Tenant}, {"memory_space_id", c.MemorySpaceID},
		{"memory_id", c.MemoryID}, {"revision_id", c.RevisionID},
		{"ingestion_event_id", c.IngestionEventID}, {"extraction_job_id", c.ExtractionJobID},
		{"context_pack_id", c.ContextPackID}, {"run_id", c.RunID},
		{"session_id", c.SessionID}, {"workspace_id", c.WorkspaceID},
		{"request_id", c.RequestID}, {"trace_id", c.TraceID},
	}
	attrs := make([]slog.Attr, 0, len(values))
	for _, value := range values {
		if value.value != "" {
			attrs = append(attrs, slog.String(value.key, value.value))
		}
	}
	return attrs
}

// WithCorrelation returns a logger carrying the non-empty correlation fields.
func WithCorrelation(logger *slog.Logger, correlation Correlation) *slog.Logger {
	attrs := correlation.Attributes()
	args := make([]any, len(attrs))
	for i := range attrs {
		args[i] = attrs[i]
	}
	return logger.With(args...)
}

// New creates a JSON logger at one of the configured service levels.
func New(output io.Writer, level string) (*slog.Logger, error) {
	if output == nil {
		return nil, errors.New("logging output is required")
	}
	var minimum slog.Level
	switch level {
	case "debug":
		minimum = slog.LevelDebug
	case "info":
		minimum = slog.LevelInfo
	case "warn":
		minimum = slog.LevelWarn
	case "error":
		minimum = slog.LevelError
	default:
		return nil, errors.New("logging level must be debug, info, warn, or error")
	}
	return slog.New(slog.NewJSONHandler(output, &slog.HandlerOptions{
		Level:       minimum,
		ReplaceAttr: redactAttribute,
	})), nil
}

// Duration emits latency in milliseconds without using a high-cardinality value.
func Duration(value time.Duration) slog.Attr {
	return slog.Int64("latency_ms", value.Milliseconds())
}

func redactAttribute(_ []string, attr slog.Attr) slog.Attr {
	key := normalizeKey(attr.Key)
	if forbiddenContentKey(key) {
		return slog.Attr{}
	}
	if secretKey(key) {
		return slog.String(attr.Key, redacted)
	}
	return attr
}

func normalizeKey(key string) string {
	return strings.ToLower(strings.NewReplacer("-", "_", ".", "_").Replace(key))
}

func forbiddenContentKey(key string) bool {
	for _, token := range []string{
		"content", "prompt", "completion", "model_output", "payload", "embedding",
		"excerpt", "message_body", "request_body", "response_body", "memory_text", "grant",
		"error", "err", "headers", "cookie", "set_cookie",
	} {
		if key == token || strings.HasSuffix(key, "_"+token) {
			return true
		}
	}
	return false
}

func secretKey(key string) bool {
	for _, token := range []string{
		"secret", "password", "passwd", "authorization", "credential", "api_key",
		"token", "access_token", "refresh_token", "bearer_token", "private_key", "database_url", "dsn",
	} {
		if key == token || strings.HasSuffix(key, "_"+token) {
			return true
		}
	}
	return false
}
