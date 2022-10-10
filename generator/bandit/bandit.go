package bandit

import (
	"database/sql"
	"time"
)

type Bandit struct {
	Nickname  string
	FirstName string
	LastName  string
	MidName   sql.NullString
	BirthDate time.Time
	Gender    bool // false - male, true - female
	Influence sql.NullInt16
}
