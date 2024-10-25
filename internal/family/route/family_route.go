package family

import (
	fRepo "pteridophyte-app/internal/family/repository"
	fHandler "pteridophyte-app/internal/family/transport/http"
	fUseCase "pteridophyte-app/internal/family/usecase"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func NewFamilyRoute(e *echo.Group, db *gorm.DB) {
	//initialize family usecase
	fRepo := fRepo.NewFamilyRepo(db)
	fUseCase := fUseCase.NewFamilyUseCase(fRepo)

	fHandler := fHandler.NewFamilyHandler(fUseCase)

	fHandler.Router(e)
}
