package model

import "time"

// Session provide session
type Session struct {
	Un           string
	LastActivity time.Time
}
