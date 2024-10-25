package usecase

import (
	entity "pteridophyte-app/internal/plant/entity"
	repository "pteridophyte-app/internal/plant/repository"
	localError "pteridophyte-app/pkg/error"

	"github.com/google/uuid"
)

type IPlantUseCase interface {
	FindAll() []entity.Plant
	Create(request entity.PlantRequestBody) (*entity.Plant, localError.GlobalError)
}

func (usecase *plantUseCase) FindAll() []entity.Plant {
	return usecase.repository.FindAll()
}

func (p *plantUseCase) Create(request entity.PlantRequestBody) (*entity.Plant, localError.GlobalError) {
	newPlant := entity.Plant{
		UUID: uuid.New(),
		Name: request.Name,
	}

	plantData, err := p.repository.Post(newPlant)
	if err != nil {
		return nil, err
	}

	return plantData, nil
}

type plantUseCase struct {
	repository repository.IPlantRepo
}

func NewPlantUseCase(repo repository.IPlantRepo) IPlantUseCase {
	return &plantUseCase{
		repository: repo,
	}
}
