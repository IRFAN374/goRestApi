package service

import (
	"context"
	"time"

	log "github.com/go-kit/kit/log"

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
			next: next,
		}
	}
}

func (M loggingMiddleware) AddGrocery(ctx context.Context) (err error) {

	defer func(begin time.Time) {
		M.logger.Log(
			"method", "AddGrocery",
			"request", logAddGroceryRequest{},
			"err", err,
			"took", time.Since(begin))

	}(time.Now())

	return M.next.AddGrocery(ctx)
}

func (M loggingMiddleware) GetGrocery(ctx context.Context) (err error) {

	defer func(begin time.Time) {
		M.logger.Log(
			"method", "GetGrocery",
			"request", logGetGroceryRequest{},
			"err", err,
			"took", time.Since(begin))

	}(time.Now())
	return M.next.GetGrocery(ctx)
}

func (M loggingMiddleware) GetAllGrocery(ctx context.Context) (err error) {
	defer func(begin time.Time) {
		M.logger.Log(
			"method", "GetAllGrocery",
			"request", logGetAllGroceryRequest{},
			"err", err,
			"took", time.Since(begin))

	}(time.Now())
	return M.next.GetAllGrocery(ctx)
}

func (M loggingMiddleware) UpdateGrocery(ctx context.Context) (err error) {
	defer func(begin time.Time) {
		M.logger.Log(
			"method", "UpdateGrocery",
			"request", logUpdateGroceryRequest{},
			"err", err,
			"took", time.Since(begin))

	}(time.Now())
	return M.next.UpdateGrocery(ctx)
}

func (M loggingMiddleware) DeleteGrocery(ctx context.Context) (err error) {
	defer func(begin time.Time) {
		M.logger.Log(
			"method", "DeleteGrocery",
			"request", logDeleteGroceryRequest{},
			"err", err,
			"took", time.Since(begin))

	}(time.Now())
	return M.next.DeleteGrocery(ctx)
}

type (
	logAddGroceryRequest    struct{}
	logGetGroceryRequest    struct{}
	logGetAllGroceryRequest struct{}
	logUpdateGroceryRequest struct{}
	logDeleteGroceryRequest struct{}
)
