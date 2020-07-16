package main

import (
	pb "client/api/proto"
	"context"
	"log"

	"google.golang.org/grpc"
)

func NewItemRepositoryClient(addr string) (pb.ItemRepositoryClient, error) {
	conn, err := grpc.Dial(addr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return nil, err
	}
	return pb.NewItemRepositoryClient(conn), nil
}

func main() {
	itemRepository, err := NewItemRepositoryClient("localhost:9094")
	if err != nil {
		log.Fatal(err)
	}

	// Test CreateItem ...
	createItemReq1 := &pb.CreateItemRequest{
		Name:  "test_item_1",
		Price: 1000,
	}
	item, err := itemRepository.CreateItem(context.Background(), createItemReq1)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("CreateItem:",*item)

	// Test CreateItem again ...
	createItemReq2 := &pb.CreateItemRequest{
		Name:  "test_item_2",
		Price: 2000,
	}
	item, err = itemRepository.CreateItem(context.Background(), createItemReq2)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("CreateItem:",*item)

	// Test UpdateItem ...
	updateItemReq := &pb.UpdateItemRequest{
		Id:    1,
		Name:  "test_item_1 updated",
		Price: 1050,
	}
	item, err = itemRepository.UpdateItem(context.Background(), updateItemReq)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("UpdateItem:",*item)

	// Test GetItem ...
	getItemReq := &pb.GetItemRequest{
		Id: 2,
	}
	item, err = itemRepository.GetItem(context.Background(), getItemReq)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("GetItem:",*item)

	// Test DeleteItem ...
	deleteItemReq := &pb.DeleteItemRequest{
		Id: 1,
	}
	item, err = itemRepository.DeleteItem(context.Background(), deleteItemReq)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("DeleteItem:",*item)

	// Try to get deleted item :-)
	getItemReq = &pb.GetItemRequest{
		Id: 1,
	}
	item, err = itemRepository.GetItem(context.Background(), getItemReq)
	if err != nil {
		log.Println("Item 1 is really deleted")
	} else {
		log.Fatal("Deleted item still exists :-/")
	}

}
