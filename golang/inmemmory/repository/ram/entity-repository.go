package ram

import (
	"sync"

	"github.com/google/uuid"
	"github.com/ic2hrmk/snips/golang/inmemmory/model"
	"github.com/ic2hrmk/snips/golang/inmemmory/repository"
)

type ramEntityRepository struct {
	storage map[string]*model.Entity
	lock    sync.Mutex
}

func NewRAMEntityRepository() repository.EntityRepository {
	return &ramEntityRepository{
		storage: make(map[string]*model.Entity),
	}
}

func (rcv *ramEntityRepository) Create(entity *model.Entity) error {
	rcv.lock.Lock()
	defer rcv.lock.Unlock()

	entity.ID = uuid.New().String()
	rcv.storage[entity.ID] = entity

	return nil
}

func (rcv *ramEntityRepository) FindByID(id string) (*model.Entity, error) {
	rcv.lock.Lock()
	defer rcv.lock.Unlock()

	entity, isFound := rcv.storage[id]
	if !isFound {
		return nil, repository.ErrEntityNotFound
	}

	return entity, nil
}
