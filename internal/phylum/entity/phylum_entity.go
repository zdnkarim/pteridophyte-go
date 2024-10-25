package phylum

import "github.com/google/uuid"

type Phylum struct {
	ID        int       `json:"-"`
	UUID      uuid.UUID `json:"uuid" gorm:"uuid"`
	KingdomID int       `json:"kingdomId" gorm:"column:kingdom_id"`
	Name      string    `json:"name" gorm:"column:name"`
	Character *string   `json:"character" gorm:"column:character"`
	Habitat   *string   `json:"habitat" gorm:"column:habitat"`
	Uses      *string   `json:"uses" gorm:"column:uses"`
	Thumbnail *string   `json:"thumbnail" gorm:"column:thumbnail"`
}

type PhylumRequestBody struct {
	Name string `json:"name" validate:"required"`
}

func (Phylum) TableName() string {
	return "master_phylum"
}
