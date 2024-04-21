package handlers

import (
	"net/http"

	"github.com/ardwiinoo/promotion-api/internal/app/models"
	"github.com/ardwiinoo/promotion-api/internal/app/services"
	"github.com/labstack/echo/v4"
)

func PSQLCreatePromotionData(PromoService services.PromotionService) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		var promo models.Promotion
		if err := ctx.Bind(&promo); err != nil {
			return echo.NewHTTPError(
				http.StatusBadRequest,
				err.Error(),
			)
		}

		newPromo, err := PromoService.CreatePromotion(promo)
		if err != nil {
			return echo.NewHTTPError(
				http.StatusInternalServerError,
				err.Error(),
			)
		}

		return ctx.JSON(
			http.StatusCreated,
			newPromo,
		)
	}
}

func PSQLGetAllPromotionData(PromoService services.PromotionService) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		promotions, err := PromoService.GetAllPromotions()
		if err != nil {
			return echo.NewHTTPError(
				http.StatusInternalServerError,
				err.Error(),
			)
		}

		return ctx.JSON(
			http.StatusOK,
			promotions,
		)
	}
}

func PSQLGetPromotionbyPromotionID(PromoService services.PromotionService) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		promotionID := ctx.Param("promotion_id")
		promotion, err := PromoService.GetPromotionbyPromotionID(promotionID)
		if err != nil {
			return echo.NewHTTPError(
				http.StatusInternalServerError,
				err.Error(),
			)
		}

		return ctx.JSON(
			http.StatusOK,
			promotion,
		)
	}
}

func PSQLUpdatePromotionbyPromotionID(PromoService services.PromotionService) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		var promo models.Promotion
		if err := ctx.Bind(&promo); err != nil {
			return echo.NewHTTPError(
				http.StatusInternalServerError,
				err.Error(),
			)
		}

		updatedPromo, err := PromoService.UpdatePromotionbyPromotionID(promo)
		if err != nil {
			return echo.NewHTTPError(
				http.StatusInternalServerError,
				err.Error(),
			)
		}

		return ctx.JSON(
			http.StatusOK,
			updatedPromo,
		)
	}
}

func PSQLDeletePromotionbyPromotionID(PromoService services.PromotionService) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		promotionID := ctx.Param("promotion_id")
		err := PromoService.DeletePromotionbyPromotionID(promotionID)
		if err != nil {
			return echo.NewHTTPError(
				http.StatusInternalServerError,
				err.Error(),
			)
		}

		return ctx.NoContent(http.StatusOK)
	}
}
