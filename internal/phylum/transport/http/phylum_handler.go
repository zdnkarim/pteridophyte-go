package phylum

import (
	entity "pteridophyte-app/internal/phylum/entity"
	"pteridophyte-app/internal/phylum/usecase"
	"pteridophyte-app/pkg/response"
	"pteridophyte-app/pkg/validation"

	"github.com/labstack/echo/v4"
)

type phylumHandler struct {
	uc usecase.IPhylumUseCase
}

func NewPhylumHandler(uc usecase.IPhylumUseCase) *phylumHandler {
	return &phylumHandler{
		uc: uc,
	}
}

func (p *phylumHandler) Router(e *echo.Group) {
	group := e.Group("/phylum")

	group.GET("", p.listAll)
	group.POST("", p.store)
}

func (p *phylumHandler) listAll(c echo.Context) error {
	phylums := p.uc.FindAll()

	// This will generate an empty response if phylums is empty
	return response.GenerateResponse(c, 200, response.WithData(phylums))
}

func (p *phylumHandler) store(c echo.Context) error {
	var request entity.PhylumRequestBody

	// Bind the request body to struct
	if err := c.Bind(&request); err != nil {
		return response.GenerateResponse(c, 400, response.WithMessage(err.Error()))
	}

	// Validate the struct
	if err := c.Validate(&request); err != nil {
		valMsg := validation.FormatValidation(err)
		return response.GenerateResponse(c, 422, response.WithMessage(valMsg.Message))
	}

	// Create Phylum
	_, errCreate := p.uc.Create(request)
	if errCreate != nil {
		return response.GenerateResponse(c, errCreate.Code, response.WithMessage(errCreate.Message))
	}

	// Response
	return response.GenerateResponse(c, 201, response.WithMessage("Phylum created"))
}
