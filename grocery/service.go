package grocery

import (
	"context"

	"github.com/IRFAN374/goRestApi/repository/mygrocery"
)

// @microgen middleware,logging
type Service interface {
	Add(ctx context.Context) (err error)
	GetGrocery(ctx context.Context) (err error)
	GetAllGrocery(ctx context.Context) (err error)

	UpdateGrocery(ctx context.Context) (err error)
	DeleteGrocery(ctx context.Context) (err error)
}

type service struct {
	groceryrepo mygrocery.Repository
}

func NewService(groceryRepo mygrocery.Repository) *service {
	return &service{
		groceryrepo: groceryRepo,
	}
}

func (s *service) Add(ctx context.Context) (err error) {

	err = s.groceryrepo.Add(ctx)
	return
}


func (s *service) GetGrocery(ctx context.Context) (err error) {

	err = s.groceryrepo.GetGrocery(ctx)
	return
}

func (s *service) GetAllGrocery(ctx context.Context) (err error) {

	err = s.groceryrepo.GetAllGrocery(ctx)
	return
}

func (s *service) UpdateGrocery(ctx context.Context) (err error) {

	err = s.groceryrepo.UpdateGrocery(ctx)
	return
}

func (s *service) DeleteGrocery(ctx context.Context) (err error) {

	err = s.groceryrepo.DeleteGrocery(ctx)
	return
}


