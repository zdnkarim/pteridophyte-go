package usecase

import (
	entity "pteridophyte-app/internal/species/entity"
	repository "pteridophyte-app/internal/species/repository"
	localError "pteridophyte-app/pkg/error"

	"github.com/google/uuid"
)

type ISpeciesUseCase interface {
	FindAll() []entity.Species
	Create(request entity.SpeciesRequestBody) (*entity.Species, localError.GlobalError)
}

func (usecase *speciesUseCase) FindAll() []entity.Species {
	return usecase.repository.FindAll()
}

func (p *speciesUseCase) Create(request entity.SpeciesRequestBody) (*entity.Species, localError.GlobalError) {
	newSpecies := entity.Species{
		UUID: uuid.New(),
		Name: request.Name,
	}

	speciesData, err := p.repository.Post(newSpecies)
	if err != nil {
		return nil, err
	}

	return speciesData, nil
}

type speciesUseCase struct {
	repository repository.ISpeciesRepo
}

func NewSpeciesUseCase(repo repository.ISpeciesRepo) ISpeciesUseCase {
	return &speciesUseCase{
		repository: repo,
	}
}
