package cli

import (
	"bytes"
	"strings"
	"testing"
)

func TestRootHelpListsPlannedCommandGroups(t *testing.T) {
	var stdout, stderr bytes.Buffer
	if code := Run(nil, &stdout, &stderr); code != 0 {
		t.Fatalf("Run() = %d, stderr = %q", code, stderr.String())
	}
	for _, group := range []string{"context", "index", "ingestion", "memory", "memory-space", "profile"} {
		if !strings.Contains(stdout.String(), "  "+group+"\n") {
			t.Errorf("help does not list %q: %s", group, stdout.String())
		}
	}
}

func TestGroupHelpListsEveryPlannedOperation(t *testing.T) {
	tests := map[string][]string{
		"memory-space": {"create", "describe"},
		"memory":       {"get", "inspect", "correct", "forget", "quarantine"},
		"context":      {"retrieve"},
		"profile":      {"inspect"},
		"ingestion":    {"status"},
		"index":        {"rebuild"},
	}
	for group, operations := range tests {
		t.Run(group, func(t *testing.T) {
			var stdout, stderr bytes.Buffer
			if code := Run([]string{group, "--help"}, &stdout, &stderr); code != 0 {
				t.Fatalf("Run() = %d, stderr = %q", code, stderr.String())
			}
			for _, operation := range operations {
				if !strings.Contains(stdout.String(), "  "+operation) {
					t.Errorf("help does not list %q: %s", operation, stdout.String())
				}
			}
		})
	}
}

func TestLeafFailsClosedUntilAPIClientExists(t *testing.T) {
	var stdout, stderr bytes.Buffer
	if code := Run([]string{"memory", "get"}, &stdout, &stderr); code != 1 {
		t.Fatalf("Run() = %d, want 1", code)
	}
	if !strings.Contains(stderr.String(), ErrNotImplemented.Error()) || !strings.Contains(stderr.String(), "public API") {
		t.Fatalf("unexpected error: %q", stderr.String())
	}
}

func TestUnknownCommandIsUsageError(t *testing.T) {
	for _, args := range [][]string{{"database", "dump"}, {"memory", "rewrite-history"}} {
		var stdout, stderr bytes.Buffer
		if code := Run(args, &stdout, &stderr); code != 2 {
			t.Fatalf("Run(%q) = %d, want 2", args, code)
		}
	}
}

func TestVersion(t *testing.T) {
	var stdout, stderr bytes.Buffer
	if code := Run([]string{"version"}, &stdout, &stderr); code != 0 {
		t.Fatalf("Run() = %d, stderr = %q", code, stderr.String())
	}
	if got, want := stdout.String(), Name+" "+Version+"\n"; got != want {
		t.Fatalf("version = %q, want %q", got, want)
	}
}
