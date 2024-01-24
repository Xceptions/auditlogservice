package models

import (
	"time"
)

type User struct {
	Username string
	Email    string
	Password string
}

type AuditInput struct {
	Actor     string      `json:"actor"`
	EventType string      `json:"action"`
	When      time.Time   `json:"when"`
	Specifics interface{} `json:"details"`
}
