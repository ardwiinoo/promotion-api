package services

import (
	"github.com/ardwiinoo/promotion-api/internal/app/models"
	"github.com/ardwiinoo/promotion-api/internal/app/repositories"
)

type PromotionServiceImpl struct {
	Repository *repositories.PromotionRepositoryImpl
}

func NewPromotionService(promotionRepo *repositories.PromotionRepositoryImpl) *PromotionServiceImpl {
	return &PromotionServiceImpl{
		Repository: promotionRepo,
	}
}

func (s *PromotionServiceImpl) CreatePromotion(promo models.Promotion) (models.Promotion, error) {
	return s.Repository.CreatePromotion(promo)
}

func (s *PromotionServiceImpl) GetAllPromotions() ([]models.Promotion, error) {
	return s.Repository.GetAllPromotions()
}

func (s *PromotionServiceImpl) GetPromotionbyPromotionID(promotionID string) (models.Promotion, error) {
	return s.Repository.GetPromotionbyPromotionID(promotionID)
}

func (s *PromotionServiceImpl) UpdatePromotionbyPromotionID(promo models.Promotion) (models.Promotion, error) {
	return s.Repository.UpdatePromotionbyPromotionID(promo)
}

func (s *PromotionServiceImpl) DeletePromotionbyPromotionID(promotionID string) error {
	return s.Repository.DeletePromotionbyPromotionID(promotionID)
}
