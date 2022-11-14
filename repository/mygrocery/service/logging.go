package service

import (
	"context"
	"time"

	log "github.com/go-kit/kit/log"

	model "github.com/IRFAN374/goRestApi/model"
	service "github.com/IRFAN374/goRestApi/repository/mygrocery"
)

type loggingMiddleware struct {
	logger log.Logger
	next   service.Repository
}

func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next service.Repository) service.Repository {
		return &loggingMiddleware{
			logger: logger,
			next:   next,
		}
	}
}

func (M loggingMiddleware) AddGrocery(ctx context.Context, grocery model.Grocery) (err error) {

	defer func(begin time.Time) {
		M.logger.Log(
			"method", "AddGrocery",
			"request", logAddGroceryRequest{
				Grocery: grocery,
			},
			"err", err,
			"took", time.Since(begin))

	}(time.Now())

	return M.next.AddGrocery(ctx, grocery)
}

func (M loggingMiddleware) GetGrocery(ctx context.Context, Name string) (grocery model.Grocery, err error) {

	defer func(begin time.Time) {
		M.logger.Log(
			"method", "GetGrocery",
			"request", logGetGroceryRequest{
				Name: Name,
			},
			"err", err,
			"took", time.Since(begin))

	}(time.Now())
	return M.next.GetGrocery(ctx, Name)
}

func (M loggingMiddleware) GetAllGrocery(ctx context.Context) (groceries []model.Grocery, err error) {
	defer func(begin time.Time) {
		M.logger.Log(
			"method", "GetAllGrocery",
			"request", logGetAllGroceryRequest{},
			"err", err,
			"took", time.Since(begin))

	}(time.Now())
	return M.next.GetAllGrocery(ctx)
}

func (M loggingMiddleware) UpdateGrocery(ctx context.Context, Name string) (grocery model.Grocery, err error) {
	defer func(begin time.Time) {
		M.logger.Log(
			"method", "UpdateGrocery",
			"request", logUpdateGroceryRequest{
				Name: Name,
			},
			"err", err,
			"took", time.Since(begin))

	}(time.Now())
	return M.next.UpdateGrocery(ctx, Name)
}

func (M loggingMiddleware) DeleteGrocery(ctx context.Context, Name string) (grocery model.Grocery, err error) {
	defer func(begin time.Time) {
		M.logger.Log(
			"method", "DeleteGrocery",
			"request", logDeleteGroceryRequest{
				Name: Name,
			},
			"err", err,
			"took", time.Since(begin))

	}(time.Now())
	return M.next.DeleteGrocery(ctx, Name)
}

type (
	logAddGroceryRequest struct {
		Grocery model.Grocery `json:"grocery"`
	}
	logGetGroceryRequest struct {
		Name string `json:"name"`
	}
	logGetAllGroceryRequest struct{}
	logUpdateGroceryRequest struct {
		Name string `json:"name"`
	}
	logDeleteGroceryRequest struct {
		Name string `json:"name"`
	}
)
