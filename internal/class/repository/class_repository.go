package class

import (
	entity "pteridophyte-app/internal/class/entity"
	localError "pteridophyte-app/pkg/error"

	"gorm.io/gorm"
)

type IClassRepo interface {
	FindAll() []entity.Class
	Post(entity entity.Class) (*entity.Class, localError.GlobalError)
}

func (repo *classRepository) FindAll() []entity.Class {
	var Class []entity.Class

	repo.db.Find(&Class)

	return Class
}

func (k *classRepository) Post(entity entity.Class) (*entity.Class, localError.GlobalError) {
	if err := k.db.Create(&entity).Error; err != nil {
		return nil, localError.ErrorInternalServer(err.Error(), err)
	}

	return &entity, nil
}

type classRepository struct {
	db *gorm.DB
}

func NewClassRepo(db *gorm.DB) IClassRepo {
	return &classRepository{
		db: db,
	}
}
