package phylum

import (
	entity "pteridophyte-app/internal/phylum/entity"
	localError "pteridophyte-app/pkg/error"

	"gorm.io/gorm"
)

type IPhylumRepo interface {
	FindAll() []entity.Phylum
	Post(entity entity.Phylum) (*entity.Phylum, localError.GlobalError)
}

func (repo *phylumRepository) FindAll() []entity.Phylum {
	var phylums []entity.Phylum

	repo.db.Find(&phylums)

	return phylums
}

func (p *phylumRepository) Post(entity entity.Phylum) (*entity.Phylum, localError.GlobalError) {
	if err := p.db.Create(&entity).Error; err != nil {
		return nil, localError.ErrorInternalServer(err.Error(), err)
	}

	return &entity, nil
}

type phylumRepository struct {
	db *gorm.DB
}

func NewPhylumRepo(db *gorm.DB) IPhylumRepo {
	return &phylumRepository{
		db: db,
	}
}
