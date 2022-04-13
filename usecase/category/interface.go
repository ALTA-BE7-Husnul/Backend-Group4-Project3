package category

import (
	_entities "project3/entities"
)

type CategoryUseCaseInterface interface {
	GetAllCategory() ([]_entities.Category, error)
}
