package dto

import "time"

// LoginRequest defines the payload for admin login.
type LoginRequest struct {
	Username  string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// OperatorInfo aggregates operator details returned to the client.
type OperatorInfo struct {
	ID       int      `json:"id"`
	Name     string   `json:"name"`
	Phone    string   `json:"phone,omitempty"`
	Email    string   `json:"email,omitempty"`
	Nickname string   `json:"nickname,omitempty"`
	Roles    []string `json:"roles,omitempty"`
	IsAdmin  bool     `json:"is_admin"`
}

// LoginResponse is the response body for a successful login.
type LoginResponse struct {
	Token        string       `json:"token"`
	ExpiresAt    time.Time    `json:"expires_at"`
	OperatorInfo OperatorInfo `json:"operator"`
}
