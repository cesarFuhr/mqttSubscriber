package fault

import "time"

type Fault struct {
	EventID string
	At      time.Time
	License string
	Fault   string
}
