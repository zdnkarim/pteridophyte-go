package plant

import (
	entity "pteridophyte-app/internal/plant/entity"
	localError "pteridophyte-app/pkg/error"

	"gorm.io/gorm"
)

type IPlantRepo interface {
	FindAll() []entity.Plant
	Post(entity entity.Plant) (*entity.Plant, localError.GlobalError)
}

func (repo *plantRepository) FindAll() []entity.Plant {
	var plants []entity.Plant

	repo.db.Find(&plants)

	return plants
}

func (p *plantRepository) Post(entity entity.Plant) (*entity.Plant, localError.GlobalError) {
	if err := p.db.Create(&entity).Error; err != nil {
		return nil, localError.ErrorInternalServer(err.Error(), err)
	}

	return &entity, nil
}

type plantRepository struct {
	db *gorm.DB
}

func NewPlantRepo(db *gorm.DB) IPlantRepo {
	return &plantRepository{
		db: db,
	}
}
