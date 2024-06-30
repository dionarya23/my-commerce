package productusecase

import (
	"encoding/json"
	"fmt"
	"time"

	"dionpamungkas.com/my-commerce/src/entities"
	"github.com/go-redis/redis"
)

func (i *sProductUsecase) FindAll(filter *entities.ProductSearchFilter) ([]*entities.Product, error) {
	cacheKey := "products"
	if filter.CategoryId != "" {
		cacheKey = fmt.Sprintf("products_%s", filter.CategoryId)
	}

	cachedProducts, err := i.redisClient.Get(cacheKey).Result()
	if err == redis.Nil {
		products, err := i.productRepository.FindAll(filter)
		if err != nil {
			return nil, err
		}

		productsJSON, err := json.Marshal(products)
		if err != nil {
			return nil, err
		}

		err = i.redisClient.Set(cacheKey, productsJSON, 10*time.Minute).Err()
		if err != nil {
			return nil, err
		}

		return products, nil
	} else if err != nil {
		return nil, err
	}

	var products []*entities.Product
	err = json.Unmarshal([]byte(cachedProducts), &products)
	if err != nil {
		return nil, err
	}

	return products, nil
}
