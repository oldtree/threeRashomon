package server

import (
	"time"
)

// real backend node info
// backend holds the info used to loadblance
// and retry status
type Backend struct {
	Name      string
	TimeStamp time.Time
}
