package models

import "gopkg.in/guregu/null.v4"

type MemoryRefreshToken struct {
	SessionID null.String
	Token     null.String
	ExpiresAt null.Time
}
