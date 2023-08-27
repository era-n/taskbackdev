package models

import "time"

type RefreshToken struct {
	Value     string    `bson:"value, omitempty"`
	CreatedAt time.Time `bson:"created_at, omitempty"`
	ExpiresAt time.Time `bson:"expires_at, omitempty"`
}
