package grocery

import (
	"context"

	"github.com/IRFAN374/goRestApi/model"
	"github.com/IRFAN374/goRestApi/repository/mygrocery"
)

// @microgen middleware,logging
type Service interface {
	AddGrocery(ctx context.Context, grocery model.Grocery) (err error)
	GetGrocery(ctx context.Context, Name string) (grocery model.Grocery, err error)
	GetAllGrocery(ctx context.Context) (groceries []model.Grocery, err error)

	UpdateGrocery(ctx context.Context, Name string) (grocery model.Grocery, err error)
	DeleteGrocery(ctx context.Context, Name string) (grocery model.Grocery, err error)
}

type service struct {
	groceryrepo mygrocery.Repository
}

func NewService(groceryRepo mygrocery.Repository) *service {
	return &service{
		groceryrepo: groceryRepo,
	}
}

func (s *service) AddGrocery(ctx context.Context, grocery model.Grocery) (err error) {
	err = s.groceryrepo.AddGrocery(ctx, grocery)
	return
}

func (s *service) GetGrocery(ctx context.Context, Name string) (grocery model.Grocery, err error) {

	grocery, err = s.groceryrepo.GetGrocery(ctx, Name)
	return
}

func (s *service) GetAllGrocery(ctx context.Context) (groceries []model.Grocery, err error) {

	groceries, err = s.groceryrepo.GetAllGrocery(ctx)
	return
}

func (s *service) UpdateGrocery(ctx context.Context, Name string) (grocery model.Grocery, err error) {

	grocery, err = s.groceryrepo.UpdateGrocery(ctx, Name)
	return
}

func (s *service) DeleteGrocery(ctx context.Context, Name string) (grocery model.Grocery, err error) {

	grocery, err = s.groceryrepo.DeleteGrocery(ctx, Name)
	return
}
