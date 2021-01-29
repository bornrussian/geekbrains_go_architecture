package main

import (
	"context"
	"log"
	"net"
	pb "server/api/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Item struct {
	ID    int32
	Name  string
	Price float32
}

type ItemRepositoryService struct {
	maxID int32
	db    map[int32]*Item
	pb.UnimplementedItemRepositoryServer
}

func (s *ItemRepositoryService) CreateItem(ctx context.Context, req *pb.CreateItemRequest) (*pb.Item, error) {
	newItem := &Item{
		Name:  req.Name,
		Price: req.Price,
	}
	s.maxID++
	newItem.ID = s.maxID

	s.db[s.maxID] = newItem

	log.Println("Created Item:", s.maxID, newItem.Name, newItem.Price)

	return &pb.Item{
		Id:    s.maxID,
		Name:  newItem.Name,
		Price: newItem.Price,
	}, nil
}

func (s *ItemRepositoryService) UpdateItem(ctx context.Context, req *pb.UpdateItemRequest) (*pb.Item, error) {
	item, ok := s.db[req.Id]
	if !ok {
		return nil, status.Error(codes.NotFound, "item not found")
	}

	item.Name = req.Name
	item.Price = req.Price

	log.Println("Updated Item:", item.ID, item.Name, item.Price)

	return &pb.Item{
		Id:    item.ID,
		Name:  item.Name,
		Price: item.Price,
	}, nil
}

func (s *ItemRepositoryService) GetItem(ctx context.Context, req *pb.GetItemRequest) (*pb.Item, error) {
	item, ok := s.db[req.Id]
	if !ok {
		return nil, status.Error(codes.NotFound, "item not found")
	}

	log.Println("Get Item:", item.ID, item.Name, item.Price)

	return &pb.Item{
		Id:    item.ID,
		Name:  item.Name,
		Price: item.Price,
	}, nil
}

func (s *ItemRepositoryService) DeleteItem(ctx context.Context, req *pb.DeleteItemRequest) (*pb.Item, error) {
	item, ok := s.db[req.Id]
	if !ok {
		return nil, status.Error(codes.NotFound, "item not found")
	}

	delete(s.db, req.Id)

	log.Println("Deleted Item:", item.ID, item.Name, item.Price)

	return &pb.Item{
		Id:    item.ID,
		Name:  item.Name,
		Price: item.Price,
	}, nil
}

func NewItemRepositoryServerStart(addr string) error {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	s := &ItemRepositoryService{
		db: make(map[int32]*Item),
	}
	serv := grpc.NewServer()

	log.Println("starting grpc server at", addr)

	pb.RegisterItemRepositoryServer(serv, s)
	if err = serv.Serve(lis); err != nil {
		return err
	}
	return nil
}

func main() {
	NewItemRepositoryServerStart("localhost:9094")
}
