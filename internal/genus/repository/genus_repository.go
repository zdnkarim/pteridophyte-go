package genus

import (
	entity "pteridophyte-app/internal/genus/entity"
	localError "pteridophyte-app/pkg/error"

	"gorm.io/gorm"
)

type IGenusRepo interface {
	FindAll() []entity.Genus
	Post(entity entity.Genus) (*entity.Genus, localError.GlobalError)
}

func (repo *genusRepository) FindAll() []entity.Genus {
	var genus []entity.Genus

	repo.db.Find(&genus)

	return genus
}

func (p *genusRepository) Post(entity entity.Genus) (*entity.Genus, localError.GlobalError) {
	if err := p.db.Create(&entity).Error; err != nil {
		return nil, localError.ErrorInternalServer(err.Error(), err)
	}

	return &entity, nil
}

type genusRepository struct {
	db *gorm.DB
}

func NewGenusRepo(db *gorm.DB) IGenusRepo {
	return &genusRepository{
		db: db,
	}
}
