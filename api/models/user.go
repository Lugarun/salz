package models

import (
  "time"
)

// Steal the name, icon from Github idp
type User struct {
  Id int `json:"id"`
  UserName string `json:"username"`
  CreatedAt time.Time `json:"created_at"`
  UpdatedAt time.Time `json:"updated_at"`
  IconPath string `json:"iconpath"`
  IdentityProvider string `json:"identityprovider"`
  IdentityProviderId string `json:"identityproviderid"`
  Elo float32 `json:"elo"`
}
