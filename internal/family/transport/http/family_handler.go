package family

import (
	entity "pteridophyte-app/internal/family/entity"
	"pteridophyte-app/internal/family/usecase"
	"pteridophyte-app/pkg/response"
	"pteridophyte-app/pkg/validation"

	"github.com/labstack/echo/v4"
)

type familyHandler struct {
	uc usecase.IFamilyUseCase
}

func NewFamilyHandler(uc usecase.IFamilyUseCase) *familyHandler {
	return &familyHandler{
		uc: uc,
	}
}

func (p *familyHandler) Router(e *echo.Group) {
	group := e.Group("/family")

	group.GET("", p.listAll)
	group.POST("", p.store)
}

func (p *familyHandler) listAll(c echo.Context) error {
	familys := p.uc.FindAll()

	// This will generate an empty response if familys is empty
	return response.GenerateResponse(c, 200, response.WithData(familys))
}

func (p *familyHandler) store(c echo.Context) error {
	var request entity.FamilyRequestBody

	// Bind the request body to struct
	if err := c.Bind(&request); err != nil {
		return response.GenerateResponse(c, 400, response.WithMessage(err.Error()))
	}

	// Validate the struct
	if err := c.Validate(&request); err != nil {
		valMsg := validation.FormatValidation(err)
		return response.GenerateResponse(c, 422, response.WithMessage(valMsg.Message))
	}

	// Create Family
	_, errCreate := p.uc.Create(request)
	if errCreate != nil {
		return response.GenerateResponse(c, errCreate.Code, response.WithMessage(errCreate.Message))
	}

	// Response
	return response.GenerateResponse(c, 201, response.WithMessage("Family created"))
}
