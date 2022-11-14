package transport

import model "github.com/IRFAN374/goRestApi/model"

type (
	AddGroceryRequest struct {
		Grocery model.Grocery `json:"grocery"`
	}
	AddGroceryResponse struct{}

	GetGroceryRequest struct {
		Name string `json:"name"`
	}
	GetGroceryResponse struct {
		Grocery model.Grocery `json:"grocery"`
	}

	GetAllGroceryRequest  struct{}
	GetAllGroceryResponse struct {
		Groceries []model.Grocery `json:"groceries"`
	}

	UpdateGroceryRequest struct {
		Name string `json:"name"`
	}
	UpdateGroceryResponse struct {
		Grocery model.Grocery `json:"grocery"`
	}

	DeleteGroceryRequest struct {
		Name string `json:"name"`
	}
	DeleteGroceryResponse struct {
		Grocery model.Grocery `json:"grocery"`
	}
)
