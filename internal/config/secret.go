package config

import (
	"encoding/json"
	"errors"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

const (
	redactedMarker    = "[REDACTED]"
	maximumSecretSize = 1 << 20
)

var environmentName = regexp.MustCompile(`^[A-Za-z_][A-Za-z0-9_]*$`)

type SecretSource string

const (
	SecretSourceEnvironment SecretSource = "environment"
	SecretSourceFile        SecretSource = "file"
)

// SecretRef is an opaque reference to operator-managed secret material.
type SecretRef struct {
	source SecretSource
	target string
}

// ParseSecretRef accepts env:NAME and file:/clean/absolute/path references.
func ParseSecretRef(value string) (SecretRef, error) {
	prefix, target, found := strings.Cut(value, ":")
	if !found || target == "" {
		return SecretRef{}, errors.New("secret reference must use an approved scheme")
	}
	switch prefix {
	case "env":
		if !environmentName.MatchString(target) {
			return SecretRef{}, errors.New("secret environment reference is invalid")
		}
		return SecretRef{source: SecretSourceEnvironment, target: target}, nil
	case "file":
		if !filepath.IsAbs(target) || filepath.Clean(target) != target {
			return SecretRef{}, errors.New("secret file reference must be a clean absolute path")
		}
		return SecretRef{source: SecretSourceFile, target: target}, nil
	default:
		return SecretRef{}, errors.New("secret reference uses an unsupported scheme")
	}
}

func (r SecretRef) IsSet() bool                  { return r.source != "" }
func (r SecretRef) Source() SecretSource         { return r.source }
func (r SecretRef) String() string               { return redactedMarker }
func (r SecretRef) GoString() string             { return redactedMarker }
func (r SecretRef) MarshalJSON() ([]byte, error) { return json.Marshal(redactedMarker) }

// Secret contains resolved secret material and redacts ordinary renderings.
type Secret struct{ value string }

func (s Secret) Value() string                { return s.value }
func (s Secret) IsSet() bool                  { return s.value != "" }
func (s Secret) String() string               { return redactedMarker }
func (s Secret) GoString() string             { return redactedMarker }
func (s Secret) MarshalJSON() ([]byte, error) { return json.Marshal(redactedMarker) }

// Resolve obtains a secret at the owning adapter boundary.
func (r SecretRef) Resolve(lookup func(string) (string, bool)) (Secret, error) {
	if !r.IsSet() {
		return Secret{}, errors.New("secret reference is not configured")
	}
	if r.source == SecretSourceEnvironment {
		value, ok := lookup(r.target)
		if !ok || value == "" {
			return Secret{}, errors.New("referenced environment secret is unavailable")
		}
		return Secret{value: value}, nil
	}
	if r.source != SecretSourceFile {
		return Secret{}, errors.New("secret reference provider is invalid")
	}
	file, err := os.Open(r.target)
	if err != nil {
		return Secret{}, errors.New("referenced file secret is unavailable")
	}
	defer func() { _ = file.Close() }()
	info, err := file.Stat()
	if err != nil || !info.Mode().IsRegular() {
		return Secret{}, errors.New("referenced file secret is not a regular file")
	}
	if info.Size() > maximumSecretSize {
		return Secret{}, errors.New("referenced file secret exceeds the size limit")
	}
	value, err := io.ReadAll(io.LimitReader(file, maximumSecretSize+1))
	if err != nil || len(value) > maximumSecretSize {
		return Secret{}, errors.New("referenced file secret cannot be read safely")
	}
	if len(value) == 0 {
		return Secret{}, errors.New("referenced file secret is empty")
	}
	return Secret{value: string(value)}, nil
}
