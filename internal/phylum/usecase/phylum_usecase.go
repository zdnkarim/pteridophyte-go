package usecase

import (
	entity "pteridophyte-app/internal/phylum/entity"
	repository "pteridophyte-app/internal/phylum/repository"
	localError "pteridophyte-app/pkg/error"

	"github.com/google/uuid"
)

type IPhylumUseCase interface {
	FindAll() []entity.Phylum
	Create(request entity.PhylumRequestBody) (*entity.Phylum, localError.GlobalError)
}

func (usecase *phylumUseCase) FindAll() []entity.Phylum {
	return usecase.repository.FindAll()
}

func (p *phylumUseCase) Create(request entity.PhylumRequestBody) (*entity.Phylum, localError.GlobalError) {
	newPhylum := entity.Phylum{
		UUID: uuid.New(),
		Name: request.Name,
	}

	phylumData, err := p.repository.Post(newPhylum)
	if err != nil {
		return nil, err
	}

	return phylumData, nil
}

type phylumUseCase struct {
	repository repository.IPhylumRepo
}

func NewPhylumUseCase(repo repository.IPhylumRepo) IPhylumUseCase {
	return &phylumUseCase{
		repository: repo,
	}
}
