package service

import (
	"context"
	"product-service/internal/app/entity"
	"product-service/internal/app/repository"
	"product-service/package/logging"
)

type CategoryService struct {
	Repo repository.CategoryRepositoryInterface
}

type CategoryServiceInterface interface {
	GetAllCategories(ctx context.Context) ([]entity.Category, error)
	CreateCategory(ctx context.Context, name string) error
}

func NewCategoryService(repo repository.CategoryRepositoryInterface) CategoryServiceInterface {
	return &CategoryService{
		Repo: repo,
	}
}

func (s *CategoryService) GetAllCategories(ctx context.Context) ([]entity.Category, error) {
	categories, err := s.Repo.GetAllCategories()
	if err != nil {
		logging.LogError(ctx, "Error GetAllCategories: %v", err)
		return nil, err
	}
	return categories, nil
}

func (s *CategoryService) CreateCategory(ctx context.Context, name string) error {

	category := &entity.Category{Name: name}

	err := s.Repo.CreateCategory(category)
	if err != nil {
		logging.LogError(ctx, "Error CreateCategory: %v", err)
		return err
	}

	return nil
}
