package repository

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	pb "shop/api/proto"

	"shop/models"
)

type Repository interface {
	CreateItem(item *models.Item) (*models.Item, error)
	GetItem(ID int32) (*models.Item, error)
	DeleteItem(ID int32) error
	UpdateItem(item *models.Item) (*models.Item, error)

	CreateOrder(Order *models.Order) (*models.Order, error)
	GetOrder(ID int32) (*models.Order, error)
}

type mapDB struct {
	ordersTable *ordersTable
	itemRepository pb.ItemRepositoryClient
}

type ordersTable struct {
	orders map[int32]*models.Order
	maxID  int32
}

func NewMapDB(cc grpc.ClientConnInterface) Repository {
	return &mapDB{
		ordersTable: &ordersTable{
			orders: make(map[int32]*models.Order),
			maxID:  0,
		},
		itemRepository: pb.NewItemRepositoryClient(cc),
	}
}

func (m *mapDB) CreateItem(item *models.Item) (*models.Item, error) {
	createItemReq := &pb.CreateItemRequest{
		Name:  item.Name,
		Price: item.Price,
	}

	newItem, err := m.itemRepository.CreateItem(context.Background(), createItemReq)
	if err != nil {
		return nil, err
	}

	return &models.Item{
		ID:    newItem.Id,
		Name:  newItem.Name,
		Price: newItem.Price,
	}, nil
}

func (m *mapDB) GetItem(ID int32) (*models.Item, error) {
	getItemReq := &pb.GetItemRequest{
		Id: ID,
	}

	item, err := m.itemRepository.GetItem(context.Background(), getItemReq)
	if err != nil {
		return nil, err
	}

	return &models.Item{
		ID:    item.Id,
		Name:  item.Name,
		Price: item.Price,
	}, nil
}

func (m *mapDB) DeleteItem(ID int32) error {
	deleteItemReq := &pb.DeleteItemRequest{
		Id: ID,
	}
	_, err := m.itemRepository.DeleteItem(context.Background(), deleteItemReq)

	if err != nil {
		return err
	}

	return nil
}

func (m *mapDB) UpdateItem(item *models.Item) (*models.Item, error) {
	updateItemReq := &pb.UpdateItemRequest{
		Id:    item.ID,
		Name:  item.Name,
		Price: item.Price,
	}
	updateItem, err := m.itemRepository.UpdateItem(context.Background(), updateItemReq)
	if err != nil {
		return nil, err
	}

	return &models.Item{
		ID:    updateItem.Id,
		Name:  updateItem.Name,
		Price: updateItem.Price,
	}, nil
}

func (m *mapDB) CreateOrder(order *models.Order) (*models.Order, error) {
	m.ordersTable.maxID++

	newOrder := &models.Order{
		ID:      m.ordersTable.maxID,
		Phone:   order.Phone,
		Email:   order.Email,
		ItemIDs: order.ItemIDs,
	}

	m.ordersTable.orders[newOrder.ID] = newOrder

	return &models.Order{
		ID:      newOrder.ID,
		Phone:   newOrder.Phone,
		Email:   newOrder.Email,
		ItemIDs: newOrder.ItemIDs,
	}, nil
}

func (m *mapDB) GetOrder(ID int32) (*models.Order, error) {
	order, ok := m.ordersTable.orders[ID]
	if !ok {
		return nil, fmt.Errorf("Order with ID: %d is not found", ID)
	}

	return &models.Order{
		ID:      order.ID,
		Phone:   order.Phone,
		ItemIDs: order.ItemIDs,
	}, nil
}
