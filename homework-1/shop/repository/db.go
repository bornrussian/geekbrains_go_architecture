package repository

import (
	"fmt"

	"shop/models"
)

type Repository interface {
	CreateItem(item *models.Item) (*models.Item, error)
	GetItem(ID int32) (*models.Item, error)
	DeleteItem(ID int32) error
	UpdateItem(item *models.Item) (*models.Item, error)
}

type mapDB struct {
	db    map[int32]*models.Item
	maxID int32
}

func NewMapDB() Repository {
	return &mapDB{
		db:    make(map[int32]*models.Item),
		maxID: 0,
	}
}

func (m *mapDB) CreateItem(item *models.Item) (*models.Item, error) {
	m.maxID++
	item.ID = m.maxID
	m.db[item.ID] = item
	return m.db[item.ID], nil
}

func (m *mapDB) GetItem(ID int32) (*models.Item, error) {
	item, ok := m.db[ID]
	if !ok {
		return nil, fmt.Errorf("Item with ID: %d is not found", ID)
	}
	return item, nil
}

func (m *mapDB) DeleteItem(ID int32) error {
	delete(m.db, ID)
	return nil
}

func (m *mapDB) UpdateItem(item *models.Item) (*models.Item, error) {
	_, ok := m.db[item.ID]
	if !ok {
		return nil, fmt.Errorf("Item with ID: %d is not found", item.ID)
	}
	m.db[item.ID] = item
	return item, nil
}
