package order

import "github.com/google/uuid"

type Order struct {
	ID        int       `json:"-"`
	UUID      uuid.UUID `json:"uuid" gorm:"uuid"`
	ClassID   int       `json:"classId" gorm:"column:class_id"`
	Name      string    `json:"name" gorm:"column:name"`
	Thumbnail *string   `json:"thumbnail" gorm:"column:thumbnail"`
}

type OrderRequestBody struct {
	Name string `json:"name" validate:"required"`
}

func (Order) TableName() string {
	return "master_order"
}
