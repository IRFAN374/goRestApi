package mygrocery

import (
	"context"
	"errors"

	model "github.com/IRFAN374/goRestApi/model"
)

var (
	ErrGroceryNameDoesNotExist = errors.New("grocery Name Does Not Exists")
)

type Repository interface {
	AddGrocery(ctx context.Context, grocery model.Grocery) (err error)
	GetGrocery(ctx context.Context, Name string) (grocery model.Grocery, err error)
	GetAllGrocery(ctx context.Context) (groceries []model.Grocery, err error)

	UpdateGrocery(ctx context.Context, Name string) (grocery model.Grocery, err error)
	DeleteGrocery(ctx context.Context, Name string) (grocery model.Grocery, err error)
}
