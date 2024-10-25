package order

import (
	oRepo "pteridophyte-app/internal/order/repository"
	oHandler "pteridophyte-app/internal/order/transport/http"
	oUseCase "pteridophyte-app/internal/order/usecase"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func NewOrderRoute(e *echo.Group, db *gorm.DB) {
	//initialize order usecase
	oRepo := oRepo.NewOrderRepo(db)
	oUseCase := oUseCase.NewOrderUseCase(oRepo)

	oHandler := oHandler.NewOrderHandler(oUseCase)

	oHandler.Router(e)
}
