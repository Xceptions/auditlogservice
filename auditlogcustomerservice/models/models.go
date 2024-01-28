package models

import (
	"time"
)

type User struct {
	Username string
	Email    string
	Password string
}

type Event struct {
	Customer  string      `json:"customer"`
	EventType string      `json:"eventtype"`
	Time      time.Time   `json:"time"`
	Specifics interface{} `json:"specifics"`
}
