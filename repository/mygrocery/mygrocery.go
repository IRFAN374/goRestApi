package mygrocery

import (
	"context"
	"errors"
)

var (
	ErrGroceryNameDoesNotExist = errors.New("grocery Name Does Not Exists")
)

type Repository interface {
	AddGrocery(ctx context.Context) (err error)
	GetGrocery(ctx context.Context) (err error)
	GetAllGrocery(ctx context.Context) (err error)

	UpdateGrocery(ctx context.Context) (err error)
	DeleteGrocery(ctx context.Context) (err error)
}
