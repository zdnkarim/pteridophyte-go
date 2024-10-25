package species

import (
	sRepo "pteridophyte-app/internal/species/repository"
	sHandler "pteridophyte-app/internal/species/transport/http"
	sUseCase "pteridophyte-app/internal/species/usecase"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func NewSpeciesRoute(e *echo.Group, db *gorm.DB) {
	//initialize species usecase
	sRepo := sRepo.NewSpeciesRepo(db)
	sUseCase := sUseCase.NewSpeciesUseCase(sRepo)

	sHandler := sHandler.NewSpeciesHandler(sUseCase)

	sHandler.Router(e)
}
