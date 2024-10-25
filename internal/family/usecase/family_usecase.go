package usecase

import (
	entity "pteridophyte-app/internal/family/entity"
	repository "pteridophyte-app/internal/family/repository"
	localError "pteridophyte-app/pkg/error"

	"github.com/google/uuid"
)

type IFamilyUseCase interface {
	FindAll() []entity.Family
	Create(request entity.FamilyRequestBody) (*entity.Family, localError.GlobalError)
}

func (usecase *familyUseCase) FindAll() []entity.Family {
	return usecase.repository.FindAll()
}

func (p *familyUseCase) Create(request entity.FamilyRequestBody) (*entity.Family, localError.GlobalError) {
	newFamily := entity.Family{
		UUID: uuid.New(),
		Name: request.Name,
	}

	familyData, err := p.repository.Post(newFamily)
	if err != nil {
		return nil, err
	}

	return familyData, nil
}

type familyUseCase struct {
	repository repository.IFamilyRepo
}

func NewFamilyUseCase(repo repository.IFamilyRepo) IFamilyUseCase {
	return &familyUseCase{
		repository: repo,
	}
}
