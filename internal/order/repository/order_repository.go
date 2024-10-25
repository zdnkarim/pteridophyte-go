package order

import (
	entity "pteridophyte-app/internal/order/entity"
	localError "pteridophyte-app/pkg/error"

	"gorm.io/gorm"
)

type IOrderRepo interface {
	FindAll() []entity.Order
	Post(entity entity.Order) (*entity.Order, localError.GlobalError)
}

func (repo *orderRepository) FindAll() []entity.Order {
	var orders []entity.Order

	repo.db.Find(&orders)

	return orders
}

func (p *orderRepository) Post(entity entity.Order) (*entity.Order, localError.GlobalError) {
	if err := p.db.Create(&entity).Error; err != nil {
		return nil, localError.ErrorInternalServer(err.Error(), err)
	}

	return &entity, nil
}

type orderRepository struct {
	db *gorm.DB
}

func NewOrderRepo(db *gorm.DB) IOrderRepo {
	return &orderRepository{
		db: db,
	}
}
