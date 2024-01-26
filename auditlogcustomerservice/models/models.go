package models

import (
	"time"
)

type User struct {
	Username string
	Email    string
	Password string
}

type AuditModel struct {
	Customer  string      `json:"actor"`
	Event     string      `json:"action"`
	Time      time.Time   `json:"when"`
	Specifics interface{} `json:"details"`
}
