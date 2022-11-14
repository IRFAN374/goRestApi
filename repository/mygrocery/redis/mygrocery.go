package mygrocery

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

const (
	userIdPrefix = "userId"
)

type repository struct {
	client *redis.ClusterClient
}

func NewRepository(client *redis.ClusterClient) *repository {
	return &repository{
		client: client,
	}
}

func (repo *repository) getKey(userId int) string {
	return fmt.Sprintf("%s:%d", userIdPrefix, userId)
}

func (repo *repository) AddGrocery(ctx context.Context) (err error) {
	key := repo.getKey(1)

	_, err = repo.client.SAdd(ctx, key, 1).Result()
	if err != nil {
		return
	}

	return
}

func (repo *repository) GetGrocery(ctx context.Context) (err error) {
	key := repo.getKey(1)

	_, err = repo.client.SAdd(ctx, key, 1).Result()
	if err != nil {
		return
	}

	return
}

func (repo *repository) GetAllGrocery(ctx context.Context) (err error) {
	key := repo.getKey(1)

	_, err = repo.client.SAdd(ctx, key, 1).Result()
	if err != nil {
		return
	}

	return
}

func (repo *repository) UpdateGrocery(ctx context.Context) (err error) {
	key := repo.getKey(1)

	_, err = repo.client.SAdd(ctx, key, 1).Result()
	if err != nil {
		return
	}

	return
}

func (repo *repository) DeleteGrocery(ctx context.Context) (err error) {
	key := repo.getKey(1)

	_, err = repo.client.SRem(ctx, key, 1).Result()
	if err != nil {
		return
	}

	return
}
