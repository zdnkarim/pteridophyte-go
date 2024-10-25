package species

import (
	entity "pteridophyte-app/internal/species/entity"
	"pteridophyte-app/internal/species/usecase"
	"pteridophyte-app/pkg/response"
	"pteridophyte-app/pkg/validation"

	"github.com/labstack/echo/v4"
)

type speciesHandler struct {
	uc usecase.ISpeciesUseCase
}

func NewSpeciesHandler(uc usecase.ISpeciesUseCase) *speciesHandler {
	return &speciesHandler{
		uc: uc,
	}
}

func (p *speciesHandler) Router(e *echo.Group) {
	group := e.Group("/species")

	group.GET("", p.listAll)
	group.POST("", p.store)
}

func (p *speciesHandler) listAll(c echo.Context) error {
	speciess := p.uc.FindAll()

	// This will generate an empty response if speciess is empty
	return response.GenerateResponse(c, 200, response.WithData(speciess))
}

func (p *speciesHandler) store(c echo.Context) error {
	var request entity.SpeciesRequestBody

	// Bind the request body to struct
	if err := c.Bind(&request); err != nil {
		return response.GenerateResponse(c, 400, response.WithMessage(err.Error()))
	}

	// Validate the struct
	if err := c.Validate(&request); err != nil {
		valMsg := validation.FormatValidation(err)
		return response.GenerateResponse(c, 422, response.WithMessage(valMsg.Message))
	}

	// Create Species
	_, errCreate := p.uc.Create(request)
	if errCreate != nil {
		return response.GenerateResponse(c, errCreate.Code, response.WithMessage(errCreate.Message))
	}

	// Response
	return response.GenerateResponse(c, 201, response.WithMessage("Species created"))
}
