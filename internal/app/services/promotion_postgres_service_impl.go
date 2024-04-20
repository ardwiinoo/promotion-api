package services

import (
	"github.com/ardwiinoo/promotion-api/internal/app/models"
	"github.com/ardwiinoo/promotion-api/internal/app/repositories"
)

type PromotionServiceImpl struct {
	Repository *repositories.PromotionRepository
}

func NewPromotionService(promotionRepo *repositories.PromotionRepository) *PromotionServiceImpl {
	return &PromotionServiceImpl{
		Repository: promotionRepo,
	}
}

func (s *PromotionServiceImpl) CreatePromotion(promo models.Promotion) (models.Promotion, error) {
	panic("implement me!")	
}

func (s *PromotionServiceImpl) GetAllPromotions() ([]models.Promotion, error) {
	return s.GetAllPromotions()
}

func (s *PromotionServiceImpl) GetPromotionbyPromotionID(promotionID string) (models.Promotion, error) {
	panic("implement me!")	
}

func (s *PromotionServiceImpl) UpdatePromotionbyPromotionID(promo models.Promotion) (models.Promotion, error) {
	panic("implement me!")	
}

func (s *PromotionServiceImpl) DeletePromotionbyPromotionID(promotionID string) error {
	panic("implement me!")	
}
