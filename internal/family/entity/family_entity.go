package family

import "github.com/google/uuid"

type Family struct {
	ID        int       `json:"-"`
	UUID      uuid.UUID `json:"uuid" gorm:"uuid"`
	OrderID   int       `json:"orderId" gorm:"column:order_id"`
	Name      string    `json:"name" gorm:"column:name"`
	Thumbnail *string   `json:"thumbnail" gorm:"column:thumbnail"`
}

type FamilyRequestBody struct {
	Name string `json:"name" validate:"required"`
}

func (Family) TableName() string {
	return "master_family"
}
