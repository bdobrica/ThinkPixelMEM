package clock

import "time"

// Clock supplies trusted server time. Application code accepts this interface
// so temporal behavior can be tested without changing the process clock.
type Clock interface {
	Now() time.Time
}

// System reads the host clock and normalizes it to UTC.
type System struct{}

func (System) Now() time.Time { return time.Now().UTC() }
