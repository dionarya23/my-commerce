package categoryusecase

import (
	"dionpamungkas.com/my-commerce/src/entities"
	category "dionpamungkas.com/my-commerce/src/repositories/category"
	"github.com/go-redis/redis"
)

type sCategoryUsecase struct {
	categoryRepository category.CategoryRepository
	redisClient        *redis.Client
}

type CategoryUsecase interface {
	FindAll() ([]*entities.Category, error)
}

func New(
	categoryRepository category.CategoryRepository,
	redisClient *redis.Client,
) CategoryUsecase {
	return &sCategoryUsecase{
		categoryRepository: categoryRepository,
		redisClient:        redisClient,
	}
}
