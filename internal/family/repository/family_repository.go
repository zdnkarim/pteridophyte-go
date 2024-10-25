package family

import (
	entity "pteridophyte-app/internal/family/entity"
	localError "pteridophyte-app/pkg/error"

	"gorm.io/gorm"
)

type IFamilyRepo interface {
	FindAll() []entity.Family
	Post(entity entity.Family) (*entity.Family, localError.GlobalError)
}

func (repo *familyRepository) FindAll() []entity.Family {
	var familys []entity.Family

	repo.db.Find(&familys)

	return familys
}

func (p *familyRepository) Post(entity entity.Family) (*entity.Family, localError.GlobalError) {
	if err := p.db.Create(&entity).Error; err != nil {
		return nil, localError.ErrorInternalServer(err.Error(), err)
	}

	return &entity, nil
}

type familyRepository struct {
	db *gorm.DB
}

func NewFamilyRepo(db *gorm.DB) IFamilyRepo {
	return &familyRepository{
		db: db,
	}
}
