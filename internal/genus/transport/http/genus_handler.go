package genus

import (
	entity "pteridophyte-app/internal/genus/entity"
	"pteridophyte-app/internal/genus/usecase"
	"pteridophyte-app/pkg/response"
	"pteridophyte-app/pkg/validation"

	"github.com/labstack/echo/v4"
)

type genusHandler struct {
	uc usecase.IGenusUseCase
}

func NewGenusHandler(uc usecase.IGenusUseCase) *genusHandler {
	return &genusHandler{
		uc: uc,
	}
}

func (p *genusHandler) Router(e *echo.Group) {
	group := e.Group("/genus")

	group.GET("", p.listAll)
	group.POST("", p.store)
}

func (p *genusHandler) listAll(c echo.Context) error {
	genuss := p.uc.FindAll()

	// This will generate an empty response if genuss is empty
	return response.GenerateResponse(c, 200, response.WithData(genuss))
}

func (p *genusHandler) store(c echo.Context) error {
	var request entity.GenusRequestBody

	// Bind the request body to struct
	if err := c.Bind(&request); err != nil {
		return response.GenerateResponse(c, 400, response.WithMessage(err.Error()))
	}

	// Validate the struct
	if err := c.Validate(&request); err != nil {
		valMsg := validation.FormatValidation(err)
		return response.GenerateResponse(c, 422, response.WithMessage(valMsg.Message))
	}

	// Create Genus
	_, errCreate := p.uc.Create(request)
	if errCreate != nil {
		return response.GenerateResponse(c, errCreate.Code, response.WithMessage(errCreate.Message))
	}

	// Response
	return response.GenerateResponse(c, 201, response.WithMessage("Genus created"))
}
