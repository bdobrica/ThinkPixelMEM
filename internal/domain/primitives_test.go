package domain

import (
	"bytes"
	"encoding/json"
	"errors"
	"testing"
	"time"
)

type fixedClock struct{ at time.Time }

func (c fixedClock) Now() time.Time { return c.at }

func TestUUIDv7RoundTripAndTypedIDs(t *testing.T) {
	wantTime := time.Date(2026, 9, 1, 12, 34, 56, 789_000_000, time.UTC)
	generator, err := NewIDGenerator(fixedClock{wantTime}, bytes.NewReader(make([]byte, 10)))
	if err != nil {
		t.Fatal(err)
	}
	id, err := generator.New()
	if err != nil {
		t.Fatal(err)
	}
	if got := id.String(); got != "01a05cf7-2895-7000-8000-000000000000" {
		t.Fatalf("UUID = %q", got)
	}
	if !id.Time().Equal(wantTime) {
		t.Fatalf("time = %s", id.Time())
	}
	parsed, err := ParseUUID(id.String())
	if err != nil || parsed != id {
		t.Fatalf("round trip = %v, %v", parsed, err)
	}
	memoryID, err := ParseMemoryID(id.String())
	if err != nil || memoryID.String() != id.String() {
		t.Fatalf("memory ID = %v, %v", memoryID, err)
	}
	encoded, err := json.Marshal(memoryID)
	if err != nil || string(encoded) != `"01a05cf7-2895-7000-8000-000000000000"` {
		t.Fatalf("JSON = %s, %v", encoded, err)
	}
	var decoded MemoryID
	if err := json.Unmarshal(encoded, &decoded); err != nil || decoded != memoryID {
		t.Fatalf("JSON round trip = %v, %v", decoded, err)
	}
}

func TestUUIDRejectsNonCanonicalOrWrongVersion(t *testing.T) {
	for _, value := range []string{
		"01990A83-3695-7000-8000-000000000000",
		"01990a83-3695-4000-8000-000000000000",
		"01990a83369570008000000000000000",
	} {
		if _, err := ParseUUID(value); err == nil {
			t.Fatalf("accepted %q", value)
		}
	}
}

func TestBoundedStringCountsRunesAndRejectsInvalidUTF8(t *testing.T) {
	value, err := NewBoundedString("mém", 3, 3)
	if err != nil || value.String() != "mém" {
		t.Fatalf("value = %q, %v", value, err)
	}
	if _, err := NewBoundedString("long", 0, 3); err == nil {
		t.Fatal("accepted overlong string")
	}
	if _, err := NewBoundedString(string([]byte{0xff}), 0, 3); err == nil {
		t.Fatal("accepted invalid UTF-8")
	}
}

func TestDigestRoundTrip(t *testing.T) {
	digest := SHA256([]byte("memory"))
	parsed, err := ParseDigest(digest.String())
	if err != nil || parsed != digest {
		t.Fatalf("round trip = %v, %v", parsed, err)
	}
	if _, err := ParseDigest("ABC"); err == nil {
		t.Fatal("accepted invalid digest")
	}
}

func TestTypedErrorPreservesCodeAndCause(t *testing.T) {
	cause := errors.New("storage")
	err := NewError(CodeConflict, "revision conflict", cause)
	if code, ok := ErrorCodeOf(err); !ok || code != CodeConflict {
		t.Fatalf("code = %q, %v", code, ok)
	}
	if !errors.Is(err, cause) {
		t.Fatal("cause was not preserved")
	}
}
