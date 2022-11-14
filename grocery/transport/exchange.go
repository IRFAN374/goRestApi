package transport

import model "github.com/IRFAN374/goRestApi/model"

type (
	AddGroceryRequest struct {
		grocery model.Grocery
	}
	AddGroceryResponse struct{}

	GetGroceryRequest struct {
		Name string
	}
	GetGroceryResponse struct {
		grocery model.Grocery
	}

	GetAllGroceryRequest  struct{}
	GetAllGroceryResponse struct {
		groceries []model.Grocery
	}

	UpdateGroceryRequest struct {
		Name string
	}
	UpdateGroceryResponse struct {
		grocery model.Grocery
	}

	DeleteGroceryRequest struct {
		Name string
	}
	DeleteGroceryResponse struct {
		grocery model.Grocery
	}
)
