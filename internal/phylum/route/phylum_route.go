package phylum

import (
	phRepo "pteridophyte-app/internal/phylum/repository"
	phHandler "pteridophyte-app/internal/phylum/transport/http"
	phUseCase "pteridophyte-app/internal/phylum/usecase"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func NewPhylumRoute(e *echo.Group, db *gorm.DB) {
	//initialize phylum usecase
	phRepo := phRepo.NewPhylumRepo(db)
	phUseCase := phUseCase.NewPhylumUseCase(phRepo)

	phHandler := phHandler.NewPhylumHandler(phUseCase)

	phHandler.Router(e)
}
