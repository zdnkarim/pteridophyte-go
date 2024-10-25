package order

import (
	entity "pteridophyte-app/internal/order/entity"
	"pteridophyte-app/internal/order/usecase"
	"pteridophyte-app/pkg/response"
	"pteridophyte-app/pkg/validation"

	"github.com/labstack/echo/v4"
)

type orderHandler struct {
	uc usecase.IOrderUseCase
}

func NewOrderHandler(uc usecase.IOrderUseCase) *orderHandler {
	return &orderHandler{
		uc: uc,
	}
}

func (p *orderHandler) Router(e *echo.Group) {
	group := e.Group("/order")

	group.GET("", p.listAll)
	group.POST("", p.store)
}

func (p *orderHandler) listAll(c echo.Context) error {
	orders := p.uc.FindAll()

	// This will generate an empty response if orders is empty
	return response.GenerateResponse(c, 200, response.WithData(orders))
}

func (p *orderHandler) store(c echo.Context) error {
	var request entity.OrderRequestBody

	// Bind the request body to struct
	if err := c.Bind(&request); err != nil {
		return response.GenerateResponse(c, 400, response.WithMessage(err.Error()))
	}

	// Validate the struct
	if err := c.Validate(&request); err != nil {
		valMsg := validation.FormatValidation(err)
		return response.GenerateResponse(c, 422, response.WithMessage(valMsg.Message))
	}

	// Create Order
	_, errCreate := p.uc.Create(request)
	if errCreate != nil {
		return response.GenerateResponse(c, errCreate.Code, response.WithMessage(errCreate.Message))
	}

	// Response
	return response.GenerateResponse(c, 201, response.WithMessage("Order created"))
}
