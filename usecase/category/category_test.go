package category

import (
	"fmt"
	_entities "project3/entities"
	"testing"

	"github.com/stretchr/testify/assert"
)


func TestGetAllCategory(t *testing.T) {
	t.Run("TestGetAllSuccess", func(t *testing.T) {
		categoryUseCase := NewCategoryUseCase(mockCategoryRepository{})
		data, err := categoryUseCase.GetAllCategory()
		assert.Nil(t, err)
		assert.Equal(t, "category 1", data[0].CategoryName)
	})

	t.Run("TestGetAllError", func(t *testing.T) {
		categoryUseCase := NewCategoryUseCase(mockCategoryRepositoryError{})
		data, err := categoryUseCase.GetAllCategory()
		assert.NotNil(t, err)
		assert.Nil(t, data)
	})
}


// === mock success ===
type mockCategoryRepository struct{}


func (m mockCategoryRepository) GetAllCategory() ([]_entities.Category, error) {
	return []_entities.Category{
		{CategoryName: "category 1"},
	}, nil
}


// === mock error ===

type mockCategoryRepositoryError struct{}

func (m mockCategoryRepositoryError) GetAllCategory() ([]_entities.Category, error) {
	return nil, fmt.Errorf("error get all data user")
}

