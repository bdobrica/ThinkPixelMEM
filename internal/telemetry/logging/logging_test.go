package logging

import (
	"bytes"
	"context"
	"encoding/json"
	"log/slog"
	"strings"
	"testing"
	"time"
)

func TestStructuredCorrelationAndLevel(t *testing.T) {
	var output bytes.Buffer
	logger, err := New(&output, "info")
	if err != nil {
		t.Fatal(err)
	}
	logger.Debug("not emitted")
	logger = WithCorrelation(logger, Correlation{
		Tenant: "tenant-hash", MemorySpaceID: "space-1", MemoryID: "memory-1",
		RunID: "run-1", WorkspaceID: "workspace-1", RequestID: "request-1", TraceID: "trace-1",
	})
	logger.InfoContext(context.Background(), "memory operation completed",
		slog.String("operation", "memory.read"), slog.String("outcome", "allowed"), Duration(1500*time.Millisecond))

	var record map[string]any
	if err := json.Unmarshal(output.Bytes(), &record); err != nil {
		t.Fatalf("decode log: %v\n%s", err, output.String())
	}
	for key, want := range map[string]any{
		"level": "INFO", "msg": "memory operation completed", "tenant": "tenant-hash",
		"memory_space_id": "space-1", "memory_id": "memory-1", "run_id": "run-1",
		"workspace_id": "workspace-1", "request_id": "request-1", "trace_id": "trace-1",
		"operation": "memory.read", "outcome": "allowed", "latency_ms": float64(1500),
	} {
		if got := record[key]; got != want {
			t.Errorf("%s = %#v, want %#v", key, got, want)
		}
	}
	if strings.Contains(output.String(), "not emitted") {
		t.Fatal("debug record was emitted below configured level")
	}
}

func TestRedactsSecretsAndDropsContent(t *testing.T) {
	const canary = "highly-sensitive-canary"
	var output bytes.Buffer
	logger, err := New(&output, "debug")
	if err != nil {
		t.Fatal(err)
	}
	logger.Info("dependency failed",
		slog.String("password", canary),
		slog.String("database.url", canary),
		slog.String("request_body", canary),
		slog.String("prompt", canary),
		slog.String("error", canary),
		slog.String("cookie", canary),
		slog.String("token", canary),
		slog.Group("dependency", slog.String("api-key", canary), slog.String("operation", "connect")),
	)
	got := output.String()
	if strings.Contains(got, canary) {
		t.Fatalf("sensitive value leaked: %s", got)
	}
	for _, forbiddenKey := range []string{"request_body", "prompt", "error", "cookie"} {
		if strings.Contains(got, forbiddenKey) {
			t.Fatalf("forbidden content field %q was logged: %s", forbiddenKey, got)
		}
	}
	if strings.Count(got, redacted) != 4 {
		t.Fatalf("expected four redacted secret fields: %s", got)
	}
}

func TestAllCanonicalCorrelationFields(t *testing.T) {
	attrs := Correlation{
		Tenant: "1", MemorySpaceID: "2", MemoryID: "3", RevisionID: "4",
		IngestionEventID: "5", ExtractionJobID: "6", ContextPackID: "7", RunID: "8",
		SessionID: "9", WorkspaceID: "10", RequestID: "11", TraceID: "12",
	}.Attributes()
	if len(attrs) != 12 {
		t.Fatalf("got %d correlation fields, want 12", len(attrs))
	}
	if got := (Correlation{RunID: "run"}).Attributes(); len(got) != 1 || got[0].Key != "run_id" {
		t.Fatalf("empty correlation fields were not omitted: %#v", got)
	}
}

func TestNewRejectsInvalidInputs(t *testing.T) {
	if _, err := New(nil, "info"); err == nil {
		t.Fatal("expected nil output to fail")
	}
	if _, err := New(&bytes.Buffer{}, "verbose"); err == nil {
		t.Fatal("expected invalid level to fail")
	}
}
