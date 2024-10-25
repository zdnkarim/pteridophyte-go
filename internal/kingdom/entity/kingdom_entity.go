package kingdom

import "github.com/google/uuid"

type Kingdom struct {
	Id        int       `json:"-"`
	Uuid      uuid.UUID `json:"uuid" gorm:"uuid"`
	Name      string    `json:"name" gorm:"column:name"`
	Thumbnail *string   `json:"thumbnail" gorm:"column:thumbnail"`
}

type KingdomRequestBody struct {
	Name      string  `json:"name" validate:"required"`
	Thumbnail *string `json:"thumbnail"`
}

type KingdomUpdateParams struct {
	Uuid string `json:"uuid" validate:"required,uuid"`
}

func (Kingdom) TableName() string {
	return "master_kingdom"
}
