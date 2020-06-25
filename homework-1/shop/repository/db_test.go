package repository

import (
	"testing"
	"shop/models"
)

func TestCreateItem(t *testing.T) {
	want := &models.Item{
		ID: 1,
		Name: "test-item-3234",
		Price: 9483,
	}

	got, err := NewMapDB().CreateItem(want)

	if err != nil {
		t.Error("CreateItem failed", err)
	}

	if got != want {
		t.Errorf("TestCreateItem failed: got: %v, want: %v", got, want)
	}
}

func TestGetItem(t *testing.T) {
	testMapDB := NewMapDB()

	sample := &models.Item{
		ID: 0,
		Name: "test-item-1239",
		Price: 7483,
	}

	want, err := testMapDB.CreateItem(sample)
	if err != nil {
		t.Error("CreateItem failed", err)
	}

	got, err := testMapDB.GetItem(want.ID)
	if err != nil {
		t.Error("GetItem failed", err)
	}

	if got != want {
		t.Errorf("TestGetItem failed: got: %v, want: %v", got, want)
	}
}

func TestDeleteItem(t *testing.T) {
	testMapDB := NewMapDB()

	sample := &models.Item{
		ID: 0,
		Name: "test-item-1239",
		Price: 7483,
	}

	sample, err := testMapDB.CreateItem(sample)
	if err != nil {
		t.Error("CreateItem failed", err)
	}

	sID := sample.ID

	err = testMapDB.DeleteItem(sID)
	if err != nil {
		t.Error("DeleteItem failed", err)
	}

	deleted, err := testMapDB.GetItem(sID)
	if err == nil {
		t.Errorf("I've suddenly got deleted item from database: %v", deleted)
	}
}

func TestUpdateItem(t *testing.T) {
	testMapDB := NewMapDB()

	originItem := &models.Item{
		ID: 0,
		Name: "test-item-1111",
		Price: 1111,
	}

	updateItem := &models.Item{
		ID: 0,
		Name: "test-item-2222",
		Price: 2222,
	}

	originItem, err := testMapDB.CreateItem(originItem)
	if err != nil {
		t.Error("CreateItem failed", err)
	}

	updateItem.ID = originItem.ID

	resultItem, err := testMapDB.UpdateItem(updateItem)
	if err != nil {
		t.Error("UpdateItem failed", err)
	}

	if resultItem != updateItem {
		t.Errorf("UpdateItem was not successfull, updateItem=%v , resultItem=%v", updateItem, resultItem)
	}
}
