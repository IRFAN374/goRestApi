package mygrocery

import (
	"context"
	"fmt"

	model "github.com/IRFAN374/goRestApi/model"
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

func (repo *repository) getKey(Name string) string {
	return fmt.Sprintf("%s:%s", userIdPrefix, Name)
}

func (repo *repository) AddGrocery(ctx context.Context, grocery model.Grocery) (err error) {
	key := repo.getKey(grocery.Name)

	_, err = repo.client.SAdd(ctx, key, grocery).Result()
	if err != nil {
		return
	}

	return
}

func (repo *repository) GetGrocery(ctx context.Context, Name string) (grocery model.Grocery, err error) {
	key := repo.getKey(Name)

	_, err = repo.client.SMembers(ctx, key).Result()
	if err != nil {
		return
	}

	return
}

func (repo *repository) GetAllGrocery(ctx context.Context) (groceries []model.Grocery, err error) {
	key := repo.getKey("Hi")

	_, err = repo.client.SAdd(ctx, key, 1).Result()
	if err != nil {
		return
	}

	return
}

func (repo *repository) UpdateGrocery(ctx context.Context, Name string) (grocery model.Grocery, err error) {
	key := repo.getKey(Name)

	_, err = repo.client.SAdd(ctx, key, 1).Result()
	if err != nil {
		return
	}

	return
}

func (repo *repository) DeleteGrocery(ctx context.Context, Name string) (grocery model.Grocery, err error) {
	key := repo.getKey(Name)

	_, err = repo.client.SRem(ctx, key, 1).Result()
	if err != nil {
		return
	}

	return
}
