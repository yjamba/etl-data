package accesslog

import "time"

type LogEntry struct {
	IP        string
	Method    string
	Status    int
	UserAgent string
	Timestamp time.Time
}
