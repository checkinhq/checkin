package domain

import "time"

type Checkin struct {
	ID     int64
	UserID int64
	Date   time.Time

	Previous     string
	GoalsReached bool
	Next         string
	Blockers     string

	CreatedAt time.Time
	UpdatedAt time.Time
}

type CheckinRepository interface {
	// Store creates or updates a checkin.
	Store(checkin *Checkin) error
}
