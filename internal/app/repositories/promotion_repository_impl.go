package repositories

import (
	"github.com/ardwiinoo/promotion-api/internal/app/models"
	"gorm.io/gorm"
)

type PromotionRepositoryImpl struct {
	db *gorm.DB
}

func NewPromotionRepository(db *gorm.DB) PromotionRepository {
	return &PromotionRepositoryImpl {
		db: db,
	}
}

func (r *PromotionRepositoryImpl) CreatePromotion(promo models.Promotion) (models.Promotion, error) {
	panic("implement me!")
}

func (r *PromotionRepositoryImpl) GetAllPromotions() ([]models.Promotion, error) {
	panic("implement me!")
}

func (r *PromotionRepositoryImpl) GetPromotionbyPromotionID(PromotionID string) (models.Promotion, error) {
	panic("implement me!")
}

func (r *PromotionRepositoryImpl) UpdatePromotionbyPromotionID(promo models.Promotion) (models.Promotion, error) {
	panic("implement me!")
}

func (r *PromotionRepositoryImpl) DeletePromotionbyPromotionID(promotionID string) error {
	panic("implement me!")
}