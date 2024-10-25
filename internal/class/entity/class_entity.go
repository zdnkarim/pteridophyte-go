package class

import "github.com/google/uuid"

type Class struct {
	ID        int       `json:"-"`
	UUID      uuid.UUID `json:"uuid" gorm:"uuid"`
	PhylumID  int       `json:"phylumId" gorm:"column:phylum_id"`
	Name      string    `json:"name" gorm:"column:name"`
	Thumbnail *string   `json:"thumbnail" gorm:"column:thumbnail"`
}

type ClassRequestBody struct {
	Name string `json:"name" validate:"required"`
}

func (Class) TableName() string {
	return "master_class"
}
