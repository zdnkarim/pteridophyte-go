package kingdom

import (
	kRepo "pteridophyte-app/internal/kingdom/repository"
	kHandler "pteridophyte-app/internal/kingdom/transport/http"
	kUseCase "pteridophyte-app/internal/kingdom/usecase"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func NewKingdomRoute(e *echo.Group, db *gorm.DB) {
	//initialize kingdom usecase
	kRepo := kRepo.NewKingdomRepo(db)
	kUseCase := kUseCase.NewKingdomUseCase(kRepo)

	kHandler := kHandler.NewKingdomHandler(kUseCase)

	kHandler.Router(e)
}
