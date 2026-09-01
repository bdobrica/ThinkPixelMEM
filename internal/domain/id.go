package domain

import (
	"crypto/rand"
	"encoding/hex"
	"io"
	"strings"
	"time"

	clockport "github.com/bdobrica/ThinkPixelMEM/internal/ports/clock"
)

// UUID is a canonical RFC 9562 UUIDv7 value.
type UUID [16]byte

// IDGenerator creates UUIDv7 values using trusted time and cryptographic
// entropy. Supplying both dependencies keeps generation deterministic in tests.
type IDGenerator struct {
	clock   clockport.Clock
	entropy io.Reader
}

func NewIDGenerator(clock clockport.Clock, entropy io.Reader) (*IDGenerator, error) {
	if clock == nil {
		return nil, NewError(CodeInternal, "ID generator clock is required", nil)
	}
	if entropy == nil {
		return nil, NewError(CodeInternal, "ID generator entropy is required", nil)
	}
	return &IDGenerator{clock: clock, entropy: entropy}, nil
}

func NewSystemIDGenerator() *IDGenerator {
	generator, _ := NewIDGenerator(clockport.System{}, rand.Reader)
	return generator
}

func (g *IDGenerator) New() (UUID, error) {
	var id UUID
	milliseconds := g.clock.Now().UnixMilli()
	if milliseconds < 0 || milliseconds > 1<<48-1 {
		return id, NewError(CodeInternal, "clock is outside UUIDv7 range", nil)
	}
	if _, err := io.ReadFull(g.entropy, id[6:]); err != nil {
		return id, NewError(CodeInternal, "generate UUIDv7 entropy", err)
	}
	for i := 5; i >= 0; i-- {
		id[i] = byte(milliseconds)
		milliseconds >>= 8
	}
	id[6] = id[6]&0x0f | 0x70
	id[8] = id[8]&0x3f | 0x80
	return id, nil
}

func ParseUUID(value string) (UUID, error) {
	var id UUID
	if len(value) != 36 || value != strings.ToLower(value) || value[8] != '-' || value[13] != '-' || value[18] != '-' || value[23] != '-' {
		return id, NewError(CodeInvalid, "ID must be a canonical lowercase UUIDv7", nil)
	}
	compact := strings.ReplaceAll(value, "-", "")
	decoded, err := hex.DecodeString(compact)
	if err != nil || len(decoded) != len(id) {
		return id, NewError(CodeInvalid, "ID must be a canonical lowercase UUIDv7", err)
	}
	copy(id[:], decoded)
	if id[6]>>4 != 7 || id[8]>>6 != 2 {
		return UUID{}, NewError(CodeInvalid, "ID must be UUIDv7", nil)
	}
	return id, nil
}

func (id UUID) String() string {
	var out [36]byte
	hex.Encode(out[0:8], id[0:4])
	out[8] = '-'
	hex.Encode(out[9:13], id[4:6])
	out[13] = '-'
	hex.Encode(out[14:18], id[6:8])
	out[18] = '-'
	hex.Encode(out[19:23], id[8:10])
	out[23] = '-'
	hex.Encode(out[24:36], id[10:16])
	return string(out[:])
}

func (id UUID) MarshalText() ([]byte, error) { return []byte(id.String()), nil }
func (id *UUID) UnmarshalText(text []byte) error {
	parsed, err := ParseUUID(string(text))
	if err == nil {
		*id = parsed
	}
	return err
}

func (id UUID) Time() time.Time {
	milliseconds := int64(id[0])<<40 | int64(id[1])<<32 | int64(id[2])<<24 | int64(id[3])<<16 | int64(id[4])<<8 | int64(id[5])
	return time.UnixMilli(milliseconds).UTC()
}

// Distinct aliases prevent identifiers from different domains being mixed.
type TenantID UUID
type PrincipalID UUID
type MemorySpaceID UUID
type MemoryID UUID
type RevisionID UUID
type EvidenceID UUID
type RunID UUID
type SessionID UUID
type WorkspaceID UUID

func (id TenantID) String() string      { return UUID(id).String() }
func (id PrincipalID) String() string   { return UUID(id).String() }
func (id MemorySpaceID) String() string { return UUID(id).String() }
func (id MemoryID) String() string      { return UUID(id).String() }
func (id RevisionID) String() string    { return UUID(id).String() }
func (id EvidenceID) String() string    { return UUID(id).String() }
func (id RunID) String() string         { return UUID(id).String() }
func (id SessionID) String() string     { return UUID(id).String() }
func (id WorkspaceID) String() string   { return UUID(id).String() }

func (id TenantID) MarshalText() ([]byte, error)      { return []byte(id.String()), nil }
func (id PrincipalID) MarshalText() ([]byte, error)   { return []byte(id.String()), nil }
func (id MemorySpaceID) MarshalText() ([]byte, error) { return []byte(id.String()), nil }
func (id MemoryID) MarshalText() ([]byte, error)      { return []byte(id.String()), nil }
func (id RevisionID) MarshalText() ([]byte, error)    { return []byte(id.String()), nil }
func (id EvidenceID) MarshalText() ([]byte, error)    { return []byte(id.String()), nil }
func (id RunID) MarshalText() ([]byte, error)         { return []byte(id.String()), nil }
func (id SessionID) MarshalText() ([]byte, error)     { return []byte(id.String()), nil }
func (id WorkspaceID) MarshalText() ([]byte, error)   { return []byte(id.String()), nil }

func (id *TenantID) UnmarshalText(text []byte) error {
	parsed, err := ParseTenantID(string(text))
	if err == nil {
		*id = parsed
	}
	return err
}
func (id *PrincipalID) UnmarshalText(text []byte) error {
	parsed, err := ParsePrincipalID(string(text))
	if err == nil {
		*id = parsed
	}
	return err
}
func (id *MemorySpaceID) UnmarshalText(text []byte) error {
	parsed, err := ParseMemorySpaceID(string(text))
	if err == nil {
		*id = parsed
	}
	return err
}
func (id *MemoryID) UnmarshalText(text []byte) error {
	parsed, err := ParseMemoryID(string(text))
	if err == nil {
		*id = parsed
	}
	return err
}
func (id *RevisionID) UnmarshalText(text []byte) error {
	parsed, err := ParseRevisionID(string(text))
	if err == nil {
		*id = parsed
	}
	return err
}
func (id *EvidenceID) UnmarshalText(text []byte) error {
	parsed, err := ParseEvidenceID(string(text))
	if err == nil {
		*id = parsed
	}
	return err
}
func (id *RunID) UnmarshalText(text []byte) error {
	parsed, err := ParseRunID(string(text))
	if err == nil {
		*id = parsed
	}
	return err
}
func (id *SessionID) UnmarshalText(text []byte) error {
	parsed, err := ParseSessionID(string(text))
	if err == nil {
		*id = parsed
	}
	return err
}
func (id *WorkspaceID) UnmarshalText(text []byte) error {
	parsed, err := ParseWorkspaceID(string(text))
	if err == nil {
		*id = parsed
	}
	return err
}

func (g *IDGenerator) NewTenantID() (TenantID, error) { id, err := g.New(); return TenantID(id), err }
func (g *IDGenerator) NewPrincipalID() (PrincipalID, error) {
	id, err := g.New()
	return PrincipalID(id), err
}
func (g *IDGenerator) NewMemorySpaceID() (MemorySpaceID, error) {
	id, err := g.New()
	return MemorySpaceID(id), err
}
func (g *IDGenerator) NewMemoryID() (MemoryID, error) { id, err := g.New(); return MemoryID(id), err }
func (g *IDGenerator) NewRevisionID() (RevisionID, error) {
	id, err := g.New()
	return RevisionID(id), err
}
func (g *IDGenerator) NewEvidenceID() (EvidenceID, error) {
	id, err := g.New()
	return EvidenceID(id), err
}
func (g *IDGenerator) NewRunID() (RunID, error) { id, err := g.New(); return RunID(id), err }
func (g *IDGenerator) NewSessionID() (SessionID, error) {
	id, err := g.New()
	return SessionID(id), err
}
func (g *IDGenerator) NewWorkspaceID() (WorkspaceID, error) {
	id, err := g.New()
	return WorkspaceID(id), err
}

func ParseTenantID(s string) (TenantID, error) { id, err := ParseUUID(s); return TenantID(id), err }
func ParsePrincipalID(s string) (PrincipalID, error) {
	id, err := ParseUUID(s)
	return PrincipalID(id), err
}
func ParseMemorySpaceID(s string) (MemorySpaceID, error) {
	id, err := ParseUUID(s)
	return MemorySpaceID(id), err
}
func ParseMemoryID(s string) (MemoryID, error) { id, err := ParseUUID(s); return MemoryID(id), err }
func ParseRevisionID(s string) (RevisionID, error) {
	id, err := ParseUUID(s)
	return RevisionID(id), err
}
func ParseEvidenceID(s string) (EvidenceID, error) {
	id, err := ParseUUID(s)
	return EvidenceID(id), err
}
func ParseRunID(s string) (RunID, error)         { id, err := ParseUUID(s); return RunID(id), err }
func ParseSessionID(s string) (SessionID, error) { id, err := ParseUUID(s); return SessionID(id), err }
func ParseWorkspaceID(s string) (WorkspaceID, error) {
	id, err := ParseUUID(s)
	return WorkspaceID(id), err
}
