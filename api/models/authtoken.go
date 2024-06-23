package models

import (
  "crypto/rand"
  "github.com/charmbracelet/log"
  b64 "encoding/base64"
  "time"
)

type AuthToken struct {
  UserId    int `json:"user_id"`
  Token     string `json:"token"`
  ExpiresAt time.Time  `json:"expires_at"`
}

func NewAuthTokenForUser(userId int) *AuthToken {
  // Generate a random token
  b := make([]byte, 32)
  _, err := rand.Read(b)
  if err != nil {
    log.Fatal("Unable to generate random token", "error", err)
  }

  b64Token := b64.StdEncoding.EncodeToString(b)

  // Expires in 1 day
  expiresAt := time.Now().Add(24 * time.Hour)

  return &AuthToken{
    UserId:    userId,
    Token:     b64Token,
    ExpiresAt: expiresAt,
  }
}
