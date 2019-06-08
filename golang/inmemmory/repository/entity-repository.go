package repository

import (
	"errors"

	"github.com/ic2hrmk/snips/golang/inmemmory/model"
)

var (
	ErrEntityNotFound = errors.New("EntityNotFound")
)

type EntityRepository interface {
	Create(entity *model.Entity) error
	FindByID(id string) (*model.Entity, error)
}
