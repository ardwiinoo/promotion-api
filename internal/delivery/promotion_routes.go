package delivery

import (
	"net/http"

	"github.com/ardwiinoo/promotion-api/internal/app/handlers"
	"github.com/ardwiinoo/promotion-api/internal/app/services"
	"github.com/labstack/echo/v4"
)

func HelloServer(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "Hello Golang!")
}

func PromotionRoute(e *echo.Echo, service services.PromotionService) {
	e.GET("/", HelloServer)
	e.GET("/api/v1/promotions", handlers.PSQLGetAllPromotionData(service))
	e.GET("/api/v1/promotions/:promotion_id", handlers.PSQLGetPromotionbyPromotionID(service))
	e.POST("/api/v1/promotions", handlers.PSQLCreatePromotionData(service))
	e.PUT("/api/v1/promotions/:promotion_id", handlers.PSQLUpdatePromotionbyPromotionID(service))
	e.DELETE("/api/v1/promotions/:promotion_id", handlers.PSQLDeletePromotionbyPromotionID(service))
}