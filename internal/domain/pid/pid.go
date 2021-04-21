package pid

import (
	"time"
)

type PID struct {
	EventID string
	At      time.Time
	License string
	PID     string
	Value   string
}
