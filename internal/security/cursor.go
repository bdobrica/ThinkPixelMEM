package security

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"strings"
	"time"

	"github.com/bdobrica/ThinkPixelMEM/internal/domain"
	clockport "github.com/bdobrica/ThinkPixelMEM/internal/ports/clock"
)

const (
	cursorVersion   = 1
	maxCursorBytes  = 2048
	maxPayloadBytes = 1024
)

// CursorCodec produces opaque, purpose-bound, expiring cursors authenticated
// with HMAC-SHA-256. A codec should receive a key from a secret provider.
type CursorCodec struct {
	key     []byte
	purpose string
	ttl     time.Duration
	clock   clockport.Clock
}

type cursorEnvelope struct {
	Version   int             `json:"v"`
	Purpose   string          `json:"p"`
	IssuedAt  int64           `json:"iat"`
	ExpiresAt int64           `json:"exp"`
	Data      json.RawMessage `json:"d"`
}

func NewCursorCodec(key []byte, purpose string, ttl time.Duration, clock clockport.Clock) (*CursorCodec, error) {
	if len(key) < 32 {
		return nil, domain.NewError(domain.CodeInternal, "cursor key must contain at least 32 bytes", nil)
	}
	if purpose == "" || len(purpose) > 100 {
		return nil, domain.NewError(domain.CodeInternal, "cursor purpose must contain 1 to 100 bytes", nil)
	}
	if ttl <= 0 || ttl > 24*time.Hour {
		return nil, domain.NewError(domain.CodeInternal, "cursor TTL must be positive and at most 24h", nil)
	}
	if clock == nil {
		return nil, domain.NewError(domain.CodeInternal, "cursor clock is required", nil)
	}
	return &CursorCodec{key: append([]byte(nil), key...), purpose: purpose, ttl: ttl, clock: clock}, nil
}

func (c *CursorCodec) Encode(value any) (string, error) {
	data, err := json.Marshal(value)
	if err != nil {
		return "", domain.NewError(domain.CodeInvalid, "cursor data is not JSON encodable", err)
	}
	if len(data) > maxPayloadBytes {
		return "", domain.NewError(domain.CodeInvalid, "cursor data exceeds 1024 bytes", nil)
	}
	now := c.clock.Now().UTC()
	payload, err := json.Marshal(cursorEnvelope{Version: cursorVersion, Purpose: c.purpose, IssuedAt: now.Unix(), ExpiresAt: now.Add(c.ttl).Unix(), Data: data})
	if err != nil {
		return "", domain.NewError(domain.CodeInternal, "encode cursor envelope", err)
	}
	signature := c.sign(payload)
	token := base64.RawURLEncoding.EncodeToString(payload) + "." + base64.RawURLEncoding.EncodeToString(signature)
	if len(token) > maxCursorBytes {
		return "", domain.NewError(domain.CodeInvalid, "cursor exceeds 2048 bytes", nil)
	}
	return token, nil
}

func (c *CursorCodec) Decode(token string, destination any) error {
	if token == "" || len(token) > maxCursorBytes || destination == nil {
		return domain.NewError(domain.CodeInvalid, "cursor is invalid", nil)
	}
	parts := strings.Split(token, ".")
	if len(parts) != 2 {
		return domain.NewError(domain.CodeInvalid, "cursor is invalid", nil)
	}
	payload, payloadErr := base64.RawURLEncoding.DecodeString(parts[0])
	signature, signatureErr := base64.RawURLEncoding.DecodeString(parts[1])
	if payloadErr != nil || signatureErr != nil || !hmac.Equal(signature, c.sign(payload)) {
		return domain.NewError(domain.CodeInvalid, "cursor authentication failed", errors.Join(payloadErr, signatureErr))
	}
	var envelope cursorEnvelope
	if err := json.Unmarshal(payload, &envelope); err != nil {
		return domain.NewError(domain.CodeInvalid, "cursor payload is invalid", err)
	}
	if envelope.Version != cursorVersion || envelope.Purpose != c.purpose || envelope.IssuedAt > envelope.ExpiresAt {
		return domain.NewError(domain.CodeInvalid, "cursor scope is invalid", nil)
	}
	now := c.clock.Now().UTC().Unix()
	if now >= envelope.ExpiresAt {
		return domain.NewError(domain.CodeExpired, "cursor has expired", nil)
	}
	if envelope.IssuedAt > now+60 {
		return domain.NewError(domain.CodeInvalid, "cursor issue time is invalid", nil)
	}
	if err := json.Unmarshal(envelope.Data, destination); err != nil {
		return domain.NewError(domain.CodeInvalid, "cursor data is invalid", err)
	}
	return nil
}

func (c *CursorCodec) sign(payload []byte) []byte {
	mac := hmac.New(sha256.New, c.key)
	_, _ = mac.Write(payload)
	return mac.Sum(nil)
}
