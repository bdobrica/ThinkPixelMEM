package security

import (
	"strings"
	"testing"
	"time"

	"github.com/bdobrica/ThinkPixelMEM/internal/domain"
)

type mutableClock struct{ at time.Time }

func (c *mutableClock) Now() time.Time { return c.at }

type pageCursor struct {
	LastID string `json:"last_id"`
	Sort   string `json:"sort"`
}

func TestCursorRoundTrip(t *testing.T) {
	clock := &mutableClock{at: time.Unix(1_800_000_000, 0)}
	codec, err := NewCursorCodec([]byte(strings.Repeat("k", 32)), "memories:list", time.Hour, clock)
	if err != nil {
		t.Fatal(err)
	}
	token, err := codec.Encode(pageCursor{LastID: "id-1", Sort: "recorded_at"})
	if err != nil {
		t.Fatal(err)
	}
	if strings.Contains(token, "id-1") {
		t.Fatal("cursor exposed plaintext state")
	}
	var got pageCursor
	if err := codec.Decode(token, &got); err != nil {
		t.Fatal(err)
	}
	if got.LastID != "id-1" || got.Sort != "recorded_at" {
		t.Fatalf("decoded = %#v", got)
	}
}

func TestCursorRejectsTamperingWrongPurposeAndExpiry(t *testing.T) {
	clock := &mutableClock{at: time.Unix(1_800_000_000, 0)}
	key := []byte(strings.Repeat("k", 32))
	codec, _ := NewCursorCodec(key, "memories:list", time.Minute, clock)
	token, _ := codec.Encode(pageCursor{LastID: "id-1"})
	var destination pageCursor
	tampered := "A" + token[1:]
	if code := errorCode(t, codec.Decode(tampered, &destination)); code != domain.CodeInvalid {
		t.Fatalf("tamper code = %q", code)
	}
	other, _ := NewCursorCodec(key, "spaces:list", time.Minute, clock)
	if code := errorCode(t, other.Decode(token, &destination)); code != domain.CodeInvalid {
		t.Fatalf("purpose code = %q", code)
	}
	clock.at = clock.at.Add(time.Minute)
	if code := errorCode(t, codec.Decode(token, &destination)); code != domain.CodeExpired {
		t.Fatalf("expiry code = %q", code)
	}
}

func TestCursorValidatesConfigurationAndPayloadSize(t *testing.T) {
	clock := &mutableClock{at: time.Now()}
	if _, err := NewCursorCodec([]byte("short"), "list", time.Hour, clock); err == nil {
		t.Fatal("accepted short key")
	}
	codec, _ := NewCursorCodec([]byte(strings.Repeat("k", 32)), "list", time.Hour, clock)
	if _, err := codec.Encode(strings.Repeat("x", maxPayloadBytes+1)); err == nil {
		t.Fatal("accepted oversized payload")
	}
}

func errorCode(t *testing.T, err error) domain.ErrorCode {
	t.Helper()
	code, ok := domain.ErrorCodeOf(err)
	if !ok {
		t.Fatalf("not a typed error: %v", err)
	}
	return code
}
