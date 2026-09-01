package domain

import (
	"crypto/sha256"
	"encoding/hex"
)

// Digest is a SHA-256 content digest. Its text form is lowercase hexadecimal.
type Digest [sha256.Size]byte

func SHA256(content []byte) Digest { return sha256.Sum256(content) }

func ParseDigest(value string) (Digest, error) {
	var out Digest
	decoded, err := hex.DecodeString(value)
	if err != nil || len(decoded) != len(out) || hex.EncodeToString(decoded) != value {
		return out, NewError(CodeInvalid, "digest must be 64 lowercase hexadecimal characters", err)
	}
	copy(out[:], decoded)
	return out, nil
}

func (d Digest) String() string               { return hex.EncodeToString(d[:]) }
func (d Digest) MarshalText() ([]byte, error) { return []byte(d.String()), nil }
