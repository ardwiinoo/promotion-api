package services

import "github.com/ardwiinoo/promotion-api/internal/app/models"

type PromotionService interface {
	CreatePromotion(promo models.Promotion) (models.Promotion, error)
	GetAllPromotions() ([]models.Promotion, error)
	GetPromotionbyPromotionID(promotionID string) (models.Promotion, error)
	UpdatePromotionbyPromotionID(promo models.Promotion) (models.Promotion, error)
	DeletePromotionbyPromotionID(promotionID string) error
}