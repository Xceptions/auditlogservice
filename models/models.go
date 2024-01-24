package models

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AuditInput struct {
	Actor   string      `json:"actor"`
	EventType  string      `json:"action"`
	When    time.Time   `json:"when"`
	Specifics interface{} `json:"details"`
}