package category

import (
	_entities "project3/entities"
)

type CategoryRepositoryInterface interface {
	GetAllCategory() ([]_entities.Category, error)
}
