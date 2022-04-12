package comment

import (
	_entities "project3/entities"

	"gorm.io/gorm"
)

type CommentRepository struct {
	DB *gorm.DB
}

func NewCommentRepository(db *gorm.DB) *CommentRepository {
	return &CommentRepository{
		DB: db,
	}
}

func (ur *CommentRepository) GetAll() ([]_entities.Comment, error) {
	var products []_entities.Comment
	tx := ur.DB.Find(&products)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return products, nil
}


func (ur *CommentRepository) CreateComment(request _entities.Comment) (_entities.Comment, error) {
	yx := ur.DB.Save(&request)
	if yx.Error != nil {
		return request, yx.Error
	}
	return request, nil
}