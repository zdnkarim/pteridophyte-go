package plant

import (
	plRepo "pteridophyte-app/internal/plant/repository"
	plHandler "pteridophyte-app/internal/plant/transport/http"
	plUseCase "pteridophyte-app/internal/plant/usecase"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func NewPlantRoute(e *echo.Group, db *gorm.DB) {
	//initialize plant usecase
	plRepo := plRepo.NewPlantRepo(db)
	plUseCase := plUseCase.NewPlantUseCase(plRepo)

	plHandler := plHandler.NewPlantHandler(plUseCase)

	plHandler.Router(e)
}
