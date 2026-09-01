package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestForbiddenPaths(t *testing.T) {
	tests := map[string]string{
		"testdata/memories/alice.json":  "test memory",
		"fixtures/test-memory.jsonl":    "test memory",
		"developer/.env.local":          "secret",
		"certificates/client.pem":       "secret",
		"qdrant_storage/raft.json":      "Qdrant",
		"qdrant/snapshots/a.snapshot":   "Qdrant",
		"data/postgres-data/PG_VERSION": "database",
		"local/state.sqlite3":           "database",
	}
	for path, want := range tests {
		if got := forbiddenPath(path); !strings.Contains(got, want) {
			t.Errorf("forbiddenPath(%q) = %q, want reason containing %q", path, got, want)
		}
	}
	for _, path := range []string{".env.example", ".env.template", "docs/memory.json", "migrations/001.sql", "docs/operations/qdrant.md"} {
		if got := forbiddenPath(path); got != "" {
			t.Errorf("forbiddenPath(%q) = %q, want allowed", path, got)
		}
	}
}

func TestScanRecognizesCredentialMaterial(t *testing.T) {
	root := t.TempDir()
	privateMarker := "-----BEGIN " + "PRIVATE KEY-----\nnot-real\n"
	modelKey := "sk-" + strings.Repeat("A", 24)
	writeTestFile(t, root, "safe.txt", "sk-example-placeholder")
	writeTestFile(t, root, "private.txt", privateMarker)
	writeTestFile(t, root, "model.txt", modelKey)

	findings, err := scan(root, []string{"safe.txt", "private.txt", "model.txt"})
	if err != nil {
		t.Fatal(err)
	}
	if len(findings) != 2 || findings[0].path != "model.txt" || findings[1].path != "private.txt" {
		t.Fatalf("unexpected findings: %#v", findings)
	}
}

func TestScanReportsForbiddenPathWithoutReadingIt(t *testing.T) {
	findings, err := scan(t.TempDir(), []string{"data/local.db"})
	if err != nil {
		t.Fatal(err)
	}
	if len(findings) != 1 || !strings.Contains(findings[0].reason, "database") {
		t.Fatalf("unexpected findings: %#v", findings)
	}
}

func writeTestFile(t *testing.T, root, name, content string) {
	t.Helper()
	path := filepath.Join(root, name)
	if err := os.WriteFile(path, []byte(content), 0o600); err != nil {
		t.Fatal(err)
	}
}
