package usecase

import (
	entity "pteridophyte-app/internal/genus/entity"
	repository "pteridophyte-app/internal/genus/repository"
	localError "pteridophyte-app/pkg/error"

	"github.com/google/uuid"
)

type IGenusUseCase interface {
	FindAll() []entity.Genus
	Create(request entity.GenusRequestBody) (*entity.Genus, localError.GlobalError)
}

func (usecase *genusUseCase) FindAll() []entity.Genus {
	return usecase.repository.FindAll()
}

func (p *genusUseCase) Create(request entity.GenusRequestBody) (*entity.Genus, localError.GlobalError) {
	newGenus := entity.Genus{
		UUID: uuid.New(),
		Name: request.Name,
	}

	genusData, err := p.repository.Post(newGenus)
	if err != nil {
		return nil, err
	}

	return genusData, nil
}

type genusUseCase struct {
	repository repository.IGenusRepo
}

func NewGenusUseCase(repo repository.IGenusRepo) IGenusUseCase {
	return &genusUseCase{
		repository: repo,
	}
}
