package genus

import (
	gRepo "pteridophyte-app/internal/genus/repository"
	gHandler "pteridophyte-app/internal/genus/transport/http"
	gUseCase "pteridophyte-app/internal/genus/usecase"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func NewGenusRoute(e *echo.Group, db *gorm.DB) {
	//initialize genus usecase
	gRepo := gRepo.NewGenusRepo(db)
	gUseCase := gUseCase.NewGenusUseCase(gRepo)

	gHandler := gHandler.NewGenusHandler(gUseCase)

	gHandler.Router(e)
}
