package inmemory

import (
	"fmt"
	"sync"

	"github.com/google/uuid"
	"github.com/slimus/fetch-task/internal/model"
)

var (
	ErrRecieptNotFound = fmt.Errorf("reciept not found")
)

type DB struct {
	reciepts map[string]*model.Reciept
	lock     sync.RWMutex
}

func NewDB() *DB {
	return &DB{
		reciepts: make(map[string]*model.Reciept),
		lock:     sync.RWMutex{},
	}
}

func (db *DB) SaveReciept(reciept *model.Reciept) (uuid.UUID, error) {
	db.lock.Lock()

	id := uuid.New()

	reciept.ID = id.String()

	db.reciepts[reciept.ID] = reciept

	db.lock.Unlock()

	//real DB can return error if something goes wrong
	return id, nil
}

func (db *DB) GetById(id string) (*model.Reciept, error) {
	reciept, ok := db.reciepts[id]
	if !ok {
		return nil, ErrRecieptNotFound
	}

	return reciept, nil
}
