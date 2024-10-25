package genus

import "github.com/google/uuid"

type Genus struct {
	ID        int       `json:"-"`
	UUID      uuid.UUID `json:"uuid" gorm:"uuid"`
	FamilyID  int       `json:"familyId" gorm:"column:family_id"`
	Name      string    `json:"name" gorm:"column:name"`
	Thumbnail *string   `json:"thumbnail" gorm:"column:thumbnail"`
}

type GenusRequestBody struct {
	Name string `json:"name" validate:"required"`
}

func (Genus) TableName() string {
	return "master_family"
}
