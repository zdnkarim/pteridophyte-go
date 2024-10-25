package usecase

import (
	entity "pteridophyte-app/internal/class/entity"
	repository "pteridophyte-app/internal/class/repository"
	localError "pteridophyte-app/pkg/error"

	"github.com/google/uuid"
)

type IClassUseCase interface {
	FindAll() []entity.Class
	Create(request entity.ClassRequestBody) (*entity.Class, localError.GlobalError)
}

func (usecase *classUseCase) FindAll() []entity.Class {
	return usecase.repository.FindAll()
}

func (p *classUseCase) Create(request entity.ClassRequestBody) (*entity.Class, localError.GlobalError) {
	newClass := entity.Class{
		UUID: uuid.New(),
		Name: request.Name,
	}

	classData, err := p.repository.Post(newClass)
	if err != nil {
		return nil, err
	}

	return classData, nil
}

type classUseCase struct {
	repository repository.IClassRepo
}

func NewClassUseCase(repo repository.IClassRepo) IClassUseCase {
	return &classUseCase{
		repository: repo,
	}
}
