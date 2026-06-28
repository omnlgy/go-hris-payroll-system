package models

import "time"

type BlacklistedToken struct {
	ID            uint      `gorm:"primaryKey" json:"id"`
	Token         string    `gorm:"uniqueIndex;not null;size:500" json:"token" validate:"required"`
	ExpiredAt     time.Time `gorm:"not null;index" json:"expired_at" validate:"required"`
	BlacklistedAt time.Time `gorm:"not null;autoCreateTime" json:"blacklisted_at"`
	CreatedAt     time.Time `json:"created_at"`
}

func (BlacklistedToken) TableName() string {
	return "blacklisted_tokens"
}