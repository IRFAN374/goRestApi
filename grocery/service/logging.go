package service

import (
	"context"
	"time"

	service "github.com/IRFAN374/goRestApi/grocery"
	log "github.com/go-kit/kit/log"
)

func LogginMiddleware(logger log.Logger) Middleware {
	return func(next service.Service) service.Service {
		return &logginMiddleware{
			logger: logger,
			next:   next,
		}
	}
}

type logginMiddleware struct {
	logger log.Logger
	next   service.Service
}

func (M logginMiddleware) AddGrocery(ctx context.Context) (err error) {

	defer func(begin time.Time) {
		M.logger.Log(
			"method", "AddGrocery",
			"request", logAddGroceryRequest{},
			"err", err,
			"took", time.Since(begin))

	}(time.Now())

	return M.next.AddGrocery(ctx)
}

func (M logginMiddleware) GetGrocery(ctx context.Context) (err error) {

	defer func(begin time.Time) {
		M.logger.Log(
			"method", "GetGrocery",
			"request", logGetGroceryRequest{},
			"err", err,
			"took", time.Since(begin))

	}(time.Now())
	return M.next.GetGrocery(ctx)
}

func (M logginMiddleware) GetAllGrocery(ctx context.Context) (err error) {
	defer func(begin time.Time) {
		M.logger.Log(
			"method", "GetAllGrocery",
			"request", logGetAllGroceryRequest{},
			"err", err,
			"took", time.Since(begin))

	}(time.Now())
	return M.next.GetAllGrocery(ctx)
}

func (M logginMiddleware) UpdateGrocery(ctx context.Context) (err error) {
	defer func(begin time.Time) {
		M.logger.Log(
			"method", "UpdateGrocery",
			"request", logUpdateGroceryRequest{},
			"err", err,
			"took", time.Since(begin))

	}(time.Now())
	return M.next.UpdateGrocery(ctx)
}

func (M logginMiddleware) DeleteGrocery(ctx context.Context) (err error) {
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
