package usecase

import (
	entity "pteridophyte-app/internal/order/entity"
	repository "pteridophyte-app/internal/order/repository"
	localError "pteridophyte-app/pkg/error"

	"github.com/google/uuid"
)

type IOrderUseCase interface {
	FindAll() []entity.Order
	Create(request entity.OrderRequestBody) (*entity.Order, localError.GlobalError)
}

func (usecase *orderUseCase) FindAll() []entity.Order {
	return usecase.repository.FindAll()
}

func (p *orderUseCase) Create(request entity.OrderRequestBody) (*entity.Order, localError.GlobalError) {
	newOrder := entity.Order{
		UUID: uuid.New(),
		Name: request.Name,
	}

	orderData, err := p.repository.Post(newOrder)
	if err != nil {
		return nil, err
	}

	return orderData, nil
}

type orderUseCase struct {
	repository repository.IOrderRepo
}

func NewOrderUseCase(repo repository.IOrderRepo) IOrderUseCase {
	return &orderUseCase{
		repository: repo,
	}
}
