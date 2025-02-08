package inmemory

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/slimus/fetch-task/internal/model"
)

var (
	ErrRecieptNotFound = fmt.Errorf("reciept not found")
)

type DB struct {
	reciepts map[string]*model.Reciept
}

func NewDB() *DB {
	return &DB{
		reciepts: make(map[string]*model.Reciept),
	}
}

func (db *DB) SaveReciept(reciept *model.Reciept) (uuid.UUID, error) {
	id := uuid.New()

	reciept.ID = id.String()

	db.reciepts[reciept.ID] = reciept

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
