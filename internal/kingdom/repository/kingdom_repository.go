package kingdom

import (
	"errors"
	entity "pteridophyte-app/internal/kingdom/entity"
	localError "pteridophyte-app/pkg/error"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type IKingdomRepo interface {
	FindAll() []entity.Kingdom
	Create(entity entity.Kingdom) (*entity.Kingdom, localError.GlobalError)
	FindByUUID(uuid uuid.UUID) (*entity.Kingdom, localError.GlobalError)
}

func (repo *kingdomRepository) FindAll() []entity.Kingdom {
	var kingdoms []entity.Kingdom

	repo.db.Find(&kingdoms)

	return kingdoms
}

func (repo *kingdomRepository) Create(entity entity.Kingdom) (*entity.Kingdom, localError.GlobalError) {
	if err := repo.db.Create(&entity).Error; err != nil {
		return nil, localError.ErrorInternalServer(err.Error(), err)
	}

	return &entity, nil
}

func (repo *kingdomRepository) FindByUUID(uuid uuid.UUID) (*entity.Kingdom, localError.GlobalError) {
	var kingdom entity.Kingdom

	if err := repo.db.Where("uuid = ?", uuid).First(&kingdom).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, localError.ErrorNotFound("Kingdom not found", err)
		}

		return nil, localError.ErrorInternalServer(err.Error(), err)
	}

	return &kingdom, nil
}

type kingdomRepository struct {
	db *gorm.DB
}

func NewKingdomRepo(db *gorm.DB) IKingdomRepo {
	return &kingdomRepository{
		db: db,
	}
}
