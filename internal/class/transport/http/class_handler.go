package class

import (
	entity "pteridophyte-app/internal/class/entity"
	"pteridophyte-app/internal/class/usecase"
	"pteridophyte-app/pkg/response"
	"pteridophyte-app/pkg/validation"

	"github.com/labstack/echo/v4"
)

type classHandler struct {
	uc usecase.IClassUseCase
}

func NewClassHandler(uc usecase.IClassUseCase) *classHandler {
	return &classHandler{
		uc: uc,
	}
}

func (p *classHandler) Router(e *echo.Group) {
	group := e.Group("/class")

	group.GET("", p.listAll)
	group.POST("", p.store)
}

func (p *classHandler) listAll(c echo.Context) error {
	class := p.uc.FindAll()

	// This will generate an empty response if class is empty
	return response.GenerateResponse(c, 200, response.WithData(class))
}

func (p *classHandler) store(c echo.Context) error {
	var request entity.ClassRequestBody

	// Bind the request body to struct
	if err := c.Bind(&request); err != nil {
		return response.GenerateResponse(c, 400, response.WithMessage(err.Error()))
	}

	// Validate the struct
	if err := c.Validate(&request); err != nil {
		valMsg := validation.FormatValidation(err)
		return response.GenerateResponse(c, 422, response.WithMessage(valMsg.Message))
	}

	// Create Class
	_, errCreate := p.uc.Create(request)
	if errCreate != nil {
		return response.GenerateResponse(c, errCreate.Code, response.WithMessage(errCreate.Message))
	}

	// Response
	return response.GenerateResponse(c, 201, response.WithMessage("Class created"))
}
