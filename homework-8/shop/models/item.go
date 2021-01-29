package models

type Item struct {
	ID    int32   `json:"id"`
	Name  string  `json:"name"`
	Price float32 `json:"price"`
}
