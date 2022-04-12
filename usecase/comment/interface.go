package comment

import (
	_entities "project3/entities"
)

type CommentUseCaseInterface interface {
	GetAll() ([]_entities.Comment, error)
	CreateComment(request _entities.Comment) (_entities.Comment, error)
}