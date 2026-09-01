// Command repositoryhygiene rejects tracked runtime data and recognizable credentials.
package main

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
)

type finding struct {
	path   string
	reason string
}

var credentialPatterns = []struct {
	reason  string
	pattern *regexp.Regexp
}{
	{"private key", regexp.MustCompile(`-----BEGIN (?:[A-Z0-9]+ )?PRIVATE KEY-----`)},
	{"Anthropic model key", regexp.MustCompile(`\bsk-ant-[A-Za-z0-9_-]{20,}\b`)},
	{"OpenAI-style model key", regexp.MustCompile(`\bsk-[A-Za-z0-9_-]{20,}\b`)},
	{"Google API key", regexp.MustCompile(`\bAIza[0-9A-Za-z_-]{35}\b`)},
	{"AWS access key", regexp.MustCompile(`\b(?:AKIA|ASIA)[A-Z0-9]{16}\b`)},
	{"Hugging Face token", regexp.MustCompile(`\bhf_[A-Za-z0-9]{20,}\b`)},
	{"GitHub token", regexp.MustCompile(`\bgh[opusr]_[A-Za-z0-9]{30,}\b`)},
}

func main() {
	root, err := repositoryRoot()
	if err != nil {
		fail(err)
	}
	paths, err := trackedPaths(root)
	if err != nil {
		fail(err)
	}
	findings, err := scan(root, paths)
	if err != nil {
		fail(err)
	}
	if len(findings) != 0 {
		fmt.Fprintln(os.Stderr, "repository hygiene check failed:")
		for _, item := range findings {
			fmt.Fprintf(os.Stderr, "  %s: %s\n", item.path, item.reason)
		}
		os.Exit(1)
	}
	fmt.Printf("Repository hygiene check passed (%d tracked files).\n", len(paths))
}

func repositoryRoot() (string, error) {
	output, err := exec.Command("git", "rev-parse", "--show-toplevel").Output()
	if err != nil {
		return "", fmt.Errorf("locate repository root: %w", err)
	}
	return strings.TrimSpace(string(output)), nil
}

func trackedPaths(root string) ([]string, error) {
	command := exec.Command("git", "ls-files", "-z")
	command.Dir = root
	output, err := command.Output()
	if err != nil {
		return nil, fmt.Errorf("list tracked files: %w", err)
	}
	items := bytes.Split(output, []byte{0})
	paths := make([]string, 0, len(items))
	for _, item := range items {
		if len(item) != 0 {
			paths = append(paths, string(item))
		}
	}
	return paths, nil
}

func scan(root string, paths []string) ([]finding, error) {
	var findings []finding
	for _, path := range paths {
		if reason := forbiddenPath(path); reason != "" {
			findings = append(findings, finding{path, reason})
			continue
		}
		content, err := os.ReadFile(filepath.Join(root, filepath.FromSlash(path)))
		if errors.Is(err, os.ErrNotExist) {
			continue // A tracked file deleted in the working tree cannot introduce data.
		}
		if err != nil {
			return nil, fmt.Errorf("read %s: %w", path, err)
		}
		if bytes.IndexByte(content, 0) >= 0 {
			continue
		}
		for _, candidate := range credentialPatterns {
			if candidate.pattern.Match(content) {
				findings = append(findings, finding{path, "contains a recognizable " + candidate.reason})
				break
			}
		}
	}
	sort.Slice(findings, func(i, j int) bool { return findings[i].path < findings[j].path })
	return findings, nil
}

func forbiddenPath(path string) string {
	normalized := strings.ToLower(filepath.ToSlash(path))
	base := filepath.Base(normalized)
	segments := "/" + normalized + "/"

	if strings.Contains(segments, "/testdata/memories/") || strings.Contains(segments, "/test/memories/") || strings.Contains(segments, "/tests/memories/") ||
		(strings.Contains(base, "test-memor") && hasExtension(base, ".json", ".jsonl", ".yaml", ".yml")) {
		return "test memory data must not be tracked"
	}
	if base == ".env" || (strings.HasPrefix(base, ".env.") && base != ".env.example" && base != ".env.template") ||
		hasExtension(base, ".pem", ".key", ".p12", ".pfx") ||
		base == "credentials.json" || base == "secrets.json" || base == "secrets.yaml" || base == "secrets.yml" {
		return "credential or secret file must not be tracked"
	}
	if strings.Contains(segments, "/qdrant_storage/") || strings.Contains(segments, "/.qdrant/") ||
		(strings.Contains(segments, "/qdrant/") && strings.Contains(segments, "/snapshots/")) || strings.HasSuffix(base, ".snapshot") {
		return "Qdrant storage or snapshot must not be tracked"
	}
	if strings.Contains(segments, "/pgdata/") || strings.Contains(segments, "/postgres-data/") ||
		hasExtension(base, ".db", ".db-shm", ".db-wal", ".sqlite", ".sqlite3") {
		return "local database artifact must not be tracked"
	}
	return ""
}

func hasExtension(name string, extensions ...string) bool {
	for _, extension := range extensions {
		if strings.HasSuffix(name, extension) {
			return true
		}
	}
	return false
}

func fail(err error) {
	fmt.Fprintln(os.Stderr, "repository hygiene check:", err)
	os.Exit(1)
}
