package usecase

import (
	"errors"
	"fmt"
	"sportscorner/pkg/domain"
	"sportscorner/pkg/repository/interfaces"
	"sportscorner/pkg/usecase/services"
)

type categoryUseCase struct {
	repository interfaces.CategoryRepository
}

func NewCategoryUseCase(repo interfaces.CategoryRepository) services.CategoryUseCase {
	return &categoryUseCase{
		repository: repo,
	}
}

func (Cat *categoryUseCase) AddCategory(category domain.Category) (domain.Category, error) {
	cat:=category.Category
	categoryExist, err := Cat.repository.CheckCategory(cat)
	if categoryExist {
		return domain.Category{}, errors.New("category alredy exist")
	}

	productResponse, err := Cat.repository.AddCategory(category)

	if err != nil {
		return domain.Category{}, err
	}

	return productResponse, nil

}

func (Cat *categoryUseCase) UpdateCategory(current string, new string) (domain.Category, error) {

	result, err := Cat.repository.CheckCategory(current)
	if err != nil {
		fmt.Println("1")
		return domain.Category{}, err
	}

	if !result {
		fmt.Println("2")
		return domain.Category{}, errors.New("there is no category as you mentioned")
	}

	newcat, err := Cat.repository.UpdateCategory(current, new)
	if err != nil {
		fmt.Println("3")
		return domain.Category{}, err
	}

	return newcat, err
}

func (Cat *categoryUseCase) DeleteCategory(categoryID string) error {

	err := Cat.repository.DeleteCategory(categoryID)
	if err != nil {
		return err
	}
	return nil

}
