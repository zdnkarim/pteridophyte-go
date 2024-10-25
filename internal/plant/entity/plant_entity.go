package plant

import "github.com/google/uuid"

type Plant struct {
	ID        int       `json:"-"`
	UUID      uuid.UUID `json:"uuid" gorm:"uuid"`
	SpeciesID int       `json:"speciesId" gorm:"column:species_id"`
	Name      string    `json:"name" gorm:"column:name"`
	Character *string   `json:"character" gorm:"column:character"`
	Habitat   *string   `json:"habitat" gorm:"column:habitat"`
	Uses      *string   `json:"uses" gorm:"column:uses"`
	Thumbnail *string   `json:"thumbnail" gorm:"column:thumbnail"`
}

type PlantRequestBody struct {
	Name string `json:"name" validate:"required"`
}

func (Plant) TableName() string {
	return "master_plant"
}
