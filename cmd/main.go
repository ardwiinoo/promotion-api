package main

import (
	"github.com/ardwiinoo/promotion-api/internal/app/repositories"
	"github.com/ardwiinoo/promotion-api/internal/app/services"
	"github.com/ardwiinoo/promotion-api/internal/config"
	"github.com/ardwiinoo/promotion-api/internal/delivery"
	"github.com/labstack/echo/v4"
)

func main() {
	config.LoadViperEnv()
	db := config.InitDatabase()

	e := echo.New()

	promoRepo := repositories.NewPromotionRepository(db)
	promoService := services.NewPromotionService(promoRepo)

	delivery.PromotionRoute(e, promoService)
	e.Logger.Fatal(
		e.Start(":8080"),
	)
}