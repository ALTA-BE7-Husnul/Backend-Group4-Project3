package comment

import (
	"fmt"
	_entities "project3/entities"
	"testing"

	"github.com/stretchr/testify/assert"
)




func TestCreateComment(t *testing.T) {
	t.Run("TestCreateCommentSuccess", func(t *testing.T) {
		commentUseCase := NewCommentUseCase(mockCommentRepository{})
		data, err := commentUseCase.CreateComment(_entities.Comment{})
		assert.Nil(t, nil, err)
		assert.Equal(t, "nice comment", data.Comment)
	})

	t.Run("TestCreateCommentSuccess", func(t *testing.T) {
		commentUseCase := NewCommentUseCase(mockCommentRepository{})
		data, err := commentUseCase.CreateComment(_entities.Comment{Comment: "nice comment"})
		assert.Nil(t, nil, err)
		assert.Equal(t, "nice comment", data.Comment)
	})

	t.Run("TestCreateCommentSuccess", func(t *testing.T) {
		commentUseCase := NewCommentUseCase(mockCommentRepository{})
		data, err := commentUseCase.CreateComment(_entities.Comment{EventID: 1})
		assert.Nil(t, nil, err)
		assert.Equal(t, uint(1), data.EventID)
	})

	t.Run("TestCreateCommentError", func(t *testing.T) {
		commentUseCase := NewCommentUseCase(mockCommentRepositoryError{})
		data, err := commentUseCase.CreateComment(_entities.Comment{})
		assert.NotNil(t, err)
		assert.Equal(t, "", data.Comment)
		assert.Nil(t, nil, err)
	})
}

func TestGetAll(t *testing.T) {
	t.Run("TestGetAllSuccess", func(t *testing.T) {
		commentUseCase := NewCommentUseCase(mockCommentRepository{})
		data, err := commentUseCase.GetAll()
		assert.Nil(t, err)
		assert.Equal(t, "nice comment", data[0].Comment)
	})

	t.Run("TestGetAllError", func(t *testing.T) {
		categoryUseCase := NewCommentUseCase(mockCommentRepositoryError{})
		data, err := categoryUseCase.GetAll()
		assert.NotNil(t, err)
		assert.Nil(t, data)
	})
}




// === mock success ===
type mockCommentRepository struct{}


func (m mockCommentRepository) CreateComment(request _entities.Comment) (_entities.Comment, error) {
	return _entities.Comment{
		EventID: 1, Comment: "nice comment",
	}, nil
}

func (m mockCommentRepository) GetAll() ([]_entities.Comment, error) {
	return []_entities.Comment{
		{EventID: 1, Comment: "nice comment"},
	}, nil
}



// === mock error ===

type mockCommentRepositoryError struct{}


func (m mockCommentRepositoryError) CreateComment(request _entities.Comment) (_entities.Comment, error) {
	return _entities.Comment{}, fmt.Errorf("error create comment")
}

func (m mockCommentRepositoryError) GetAll() ([]_entities.Comment, error) {
	return nil, fmt.Errorf("error get all comment")
}



