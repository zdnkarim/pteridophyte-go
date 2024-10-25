package class

import (
	cRepo "pteridophyte-app/internal/class/repository"
	cHandler "pteridophyte-app/internal/class/transport/http"
	cUseCase "pteridophyte-app/internal/class/usecase"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func NewClassRoute(e *echo.Group, db *gorm.DB) {
	//initialize class usecase
	cRepo := cRepo.NewClassRepo(db)
	cUseCase := cUseCase.NewClassUseCase(cRepo)

	cHandler := cHandler.NewClassHandler(cUseCase)

	cHandler.Router(e)
}
