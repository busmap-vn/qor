package auth_identity

import (
  "time"

  "github.com/jinzhu/gorm"
  "github.com/busmap-vn/qor/auth/claims"
)

// AuthIdentity auth identity session model
type AuthIdentity struct {
  gorm.Model
  Basic
  SignLogs
}

func (AuthIdentity) TableName() string { return "basics" }

// Basic basic information about auth identity
type Basic struct {
  Provider          string // phone, email, wechat, github...
  UID               string `gorm:"column:uid"`
  EncryptedPassword string
  UserID            string
  ConfirmedAt       *time.Time
}

// ToClaims convert to auth Claims
func (basic Basic) ToClaims() *claims.Claims {
  claims := claims.Claims{}
  claims.Provider = basic.Provider
  claims.Id = basic.UID
  claims.UserID = basic.UserID
  return &claims
}
