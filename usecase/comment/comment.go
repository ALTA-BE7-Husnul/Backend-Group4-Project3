package comment

import (
	"errors"
	_entities "project3/entities"
	_commentRepository "project3/repository/comment"
)

type CommentUseCase struct {
	commentRepository _commentRepository.CommentRepositoryInterface
}

func NewCommentUseCase(commentRepo _commentRepository.CommentRepositoryInterface) CommentUseCaseInterface {
	return &CommentUseCase{
		commentRepository: commentRepo,
	}
}

func (uuc *CommentUseCase) GetAll() ([]_entities.Comment, error) {
	products, err := uuc.commentRepository.GetAll()
	return products, err
}

func (uuc *CommentUseCase) CreateComment(request _entities.Comment) (_entities.Comment, error) {
	comment, err := uuc.commentRepository.CreateComment(request)
	if request.EventID == 0 {
		return comment, errors.New("can't be empty")
	}
	if request.Comment == "" {
		return comment, errors.New("can't be empty")
	}
	
	return comment, err
}