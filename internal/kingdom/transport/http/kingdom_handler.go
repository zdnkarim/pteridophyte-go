package kingdom

import (
	entity "pteridophyte-app/internal/kingdom/entity"
	"pteridophyte-app/internal/kingdom/usecase"
	"pteridophyte-app/pkg/response"
	"pteridophyte-app/pkg/validation"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type kingdomHandler struct {
	uc usecase.IKingdomUseCase
}

func NewKingdomHandler(uc usecase.IKingdomUseCase) *kingdomHandler {
	return &kingdomHandler{
		uc: uc,
	}
}

func (k *kingdomHandler) Router(e *echo.Group) {
	group := e.Group("/kingdom")

	group.GET("", k.findAll)
	group.POST("", k.create)
	group.GET("/:uuid", k.FindByUUID)
}

func (k *kingdomHandler) findAll(c echo.Context) error {
	kingdoms := k.uc.FindAll()

	// This will generate an empty response if kingdoms is empty
	return response.GenerateResponse(c, 200, response.WithData(kingdoms))
}

func (k *kingdomHandler) create(c echo.Context) error {
	var request entity.KingdomRequestBody

	// Bind the request body to struct
	if err := c.Bind(&request); err != nil {
		return response.GenerateResponse(c, 400, response.WithMessage(err.Error()))
	}

	// Validate the struct
	if err := c.Validate(&request); err != nil {
		valMsg := validation.FormatValidation(err)
		return response.GenerateResponse(c, 422, response.WithMessage(valMsg.Message))
	}

	// Create Kingdom
	_, errCreate := k.uc.Create(request)
	if errCreate != nil {
		return response.GenerateResponse(c, errCreate.Code, response.WithMessage(errCreate.Message))
	}

	// Response
	return response.GenerateResponse(c, 201, response.WithMessage("Kingdom created"))
}

func (k *kingdomHandler) FindByUUID(c echo.Context) error {
	var key entity.KingdomUpdateParams

	key.Uuid = c.Param("uuid")

	if err := c.Validate(&key); err != nil {
		valMsg := validation.FormatValidation(err)
		return response.GenerateResponse(c, 422, response.WithMessage(valMsg.Message))
	}

	kingdom, err := k.uc.FindByUUID(uuid.MustParse(key.Uuid))
	if err != nil {
		return response.GenerateResponse(c, err.Code, response.WithMessage(err.Message))
	}

	return response.GenerateResponse(c, 200, response.WithData(kingdom))
}
