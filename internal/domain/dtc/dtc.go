package dtc

import "time"

type DTC struct {
	EventID     string
	At          time.Time
	License     string
	DTC         string
	Description string
}
