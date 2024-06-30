package productusecase

import (
	"dionpamungkas.com/my-commerce/src/entities"
	product "dionpamungkas.com/my-commerce/src/repositories/product"
	"github.com/go-redis/redis"
)

type sProductUsecase struct {
	productRepository product.ProductRepository
	redisClient       *redis.Client
}

type ProductUsecase interface {
	FindAll(filter *entities.ProductSearchFilter) ([]*entities.Product, error)
}

func New(
	productRepository product.ProductRepository,
	redisClient *redis.Client,
) ProductUsecase {
	return &sProductUsecase{
		productRepository: productRepository,
		redisClient:       redisClient,
	}
}
