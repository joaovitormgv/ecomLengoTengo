package models

import "time"

type Session struct {
	SessionID string    `json:"session_id"`
	Data      []byte    `json:"data"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type SessionNavigationHistory struct {
	ID          int       `json:"id"`
	SessionID   string    `json:"session_id"`
	ProductID   int       `json:"product_id"`
	TimeVisited time.Time `json:"time_visited"`
	ActionTaken string    `json:"action_taken"`
}
