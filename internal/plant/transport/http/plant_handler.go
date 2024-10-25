package plant

import (
	entity "pteridophyte-app/internal/plant/entity"
	"pteridophyte-app/internal/plant/usecase"
	"pteridophyte-app/pkg/response"
	"pteridophyte-app/pkg/validation"

	"github.com/labstack/echo/v4"
)

type plantHandler struct {
	uc usecase.IPlantUseCase
}

func NewPlantHandler(uc usecase.IPlantUseCase) *plantHandler {
	return &plantHandler{
		uc: uc,
	}
}

func (p *plantHandler) Router(e *echo.Group) {
	group := e.Group("/plant")

	group.GET("", p.listAll)
	group.POST("", p.store)
}

func (p *plantHandler) listAll(c echo.Context) error {
	plants := p.uc.FindAll()

	// This will generate an empty response if plants is empty
	return response.GenerateResponse(c, 200, response.WithData(plants))
}

func (p *plantHandler) store(c echo.Context) error {
	var request entity.PlantRequestBody

	// Bind the request body to struct
	if err := c.Bind(&request); err != nil {
		return response.GenerateResponse(c, 400, response.WithMessage(err.Error()))
	}

	// Validate the struct
	if err := c.Validate(&request); err != nil {
		valMsg := validation.FormatValidation(err)
		return response.GenerateResponse(c, 422, response.WithMessage(valMsg.Message))
	}

	// Create Plant
	_, errCreate := p.uc.Create(request)
	if errCreate != nil {
		return response.GenerateResponse(c, errCreate.Code, response.WithMessage(errCreate.Message))
	}

	// Response
	return response.GenerateResponse(c, 201, response.WithMessage("Plant created"))
}
