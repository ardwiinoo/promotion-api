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
	
	result := r.db.Create(&promo)
	if result.Error != nil {
		return models.Promotion{}, result.Error
	}

	return promo, nil
}

func (r *PromotionRepositoryImpl) GetAllPromotions() ([]models.Promotion, error) {
	
	var promotions []models.Promotion
	result := r.db.Find(&promotions)
	if result.Error != nil {
		return nil, result.Error
	}

	return promotions, nil
}

func (r *PromotionRepositoryImpl) GetPromotionbyPromotionID(promotionID string) (models.Promotion, error) {
	
	var promotion models.Promotion
	result := r.db.Where("promotion_id = ?", promotionID).First(&promotion) 
	if result.Error != nil {
		return models.Promotion{}, result.Error
	}

	return promotion, nil
}

func (r *PromotionRepositoryImpl) UpdatePromotionbyPromotionID(promo models.Promotion) (models.Promotion, error) {
	
	result := r.db.Model(&promo).Where("promotion_id = ?", promo.PromotionID).Updates(promo)
	if result.Error != nil {
		return models.Promotion{}, result.Error
	}
	return promo, nil
}

func (r *PromotionRepositoryImpl) DeletePromotionbyPromotionID(promotionID string) error {
	
	var promotion models.Promotion
	result := r.db.Where("promotion_id = ?", promotionID).Delete(&promotion)
	return result.Error
}