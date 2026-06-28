package repository

import (
	"time"

	"github.com/omnlgy/go-hris-payroll-system/internal/models"
	"gorm.io/gorm"
)

var BlacklistedTokenNotFound = gorm.ErrRecordNotFound

type BlacklistedTokenRepository struct {
	db *gorm.DB
}

func NewBlacklistedTokenRepository(db *gorm.DB) *BlacklistedTokenRepository {
	return &BlacklistedTokenRepository{
		db: db,
	}
}

func (r *BlacklistedTokenRepository) Blacklist(token string, expiredAt time.Time) (*models.BlacklistedToken, error) {
	bt := &models.BlacklistedToken{
		Token:         token,
		ExpiredAt:     expiredAt,
		BlacklistedAt: time.Now(),
	}
	if err := r.db.Create(bt).Error; err != nil {
		return nil, err
	}
	return bt, nil
}

func (r *BlacklistedTokenRepository) IsBlacklisted(token string) (bool, error) {
	var count int64
	err := r.db.Model(&models.BlacklistedToken{}).
		Where("token = ? AND expired_at > NOW()", token).
		Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (r *BlacklistedTokenRepository) CleanupExpired() error {
	return r.db.Where("expired_at <= NOW()").Delete(&models.BlacklistedToken{}).Error
}