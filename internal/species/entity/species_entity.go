package species

import "github.com/google/uuid"

type Species struct {
	ID        int       `json:"-"`
	UUID      uuid.UUID `json:"uuid" gorm:"uuid"`
	SpeciesID int       `json:"genusId" gorm:"column:genus_id"`
	Name      string    `json:"name" gorm:"column:name"`
	Thumbnail *string   `json:"thumbnail" gorm:"column:thumbnail"`
}

type SpeciesRequestBody struct {
	Name string `json:"name" validate:"required"`
}

func (Species) TableName() string {
	return "master_species"
}
