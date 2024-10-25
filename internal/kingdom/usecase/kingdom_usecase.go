package usecase

import (
	entity "pteridophyte-app/internal/kingdom/entity"
	repository "pteridophyte-app/internal/kingdom/repository"
	localError "pteridophyte-app/pkg/error"

	"github.com/google/uuid"
)

type IKingdomUseCase interface {
	FindAll() []entity.Kingdom
	Create(entity entity.KingdomRequestBody) (*entity.Kingdom, localError.GlobalError)
	FindByUUID(uuid uuid.UUID) (*entity.Kingdom, localError.GlobalError)
}

func (usecase *kingdomUseCase) FindAll() []entity.Kingdom {
	return usecase.repository.FindAll()
}

func (usecase *kingdomUseCase) Create(request entity.KingdomRequestBody) (*entity.Kingdom, localError.GlobalError) {
	newKingdom := entity.Kingdom{
		Uuid:      uuid.New(),
		Name:      request.Name,
		Thumbnail: request.Thumbnail, // Directly assign, nil will be stored as null in DB
	}

	kingdomData, err := usecase.repository.Create(newKingdom)
	if err != nil {
		return nil, err
	}

	return kingdomData, nil
}

func (usecase *kingdomUseCase) FindByUUID(uuid uuid.UUID) (*entity.Kingdom, localError.GlobalError) {
	return usecase.repository.FindByUUID(uuid)
}

type kingdomUseCase struct {
	repository repository.IKingdomRepo
}

func NewKingdomUseCase(repo repository.IKingdomRepo) IKingdomUseCase {
	return &kingdomUseCase{
		repository: repo,
	}
}
