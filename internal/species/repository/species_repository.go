package species

import (
	entity "pteridophyte-app/internal/species/entity"
	localError "pteridophyte-app/pkg/error"

	"gorm.io/gorm"
)

type ISpeciesRepo interface {
	FindAll() []entity.Species
	Post(entity entity.Species) (*entity.Species, localError.GlobalError)
}

func (repo *speciesRepository) FindAll() []entity.Species {
	var species []entity.Species

	repo.db.Find(&species)

	return species
}

func (p *speciesRepository) Post(entity entity.Species) (*entity.Species, localError.GlobalError) {
	if err := p.db.Create(&entity).Error; err != nil {
		return nil, localError.ErrorInternalServer(err.Error(), err)
	}

	return &entity, nil
}

type speciesRepository struct {
	db *gorm.DB
}

func NewSpeciesRepo(db *gorm.DB) ISpeciesRepo {
	return &speciesRepository{
		db: db,
	}
}
