package categoryusecase

import (
	"encoding/json"
	"time"

	"dionpamungkas.com/my-commerce/src/entities"
	"github.com/go-redis/redis"
)

func (i *sCategoryUsecase) FindAll() ([]*entities.Category, error) {
	cacheKey := "categories"

	cachedCategories, err := i.redisClient.Get(cacheKey).Result()
	if err == redis.Nil {
		categories, err := i.categoryRepository.FindAll()
		if err != nil {
			return nil, err
		}

		categoriesJSON, err := json.Marshal(categories)
		if err != nil {
			return nil, err
		}

		err = i.redisClient.Set(cacheKey, categoriesJSON, 10*time.Minute).Err()
		if err != nil {
			return nil, err
		}

		return categories, nil
	} else if err != nil {
		return nil, err
	}

	var categories []*entities.Category
	err = json.Unmarshal([]byte(cachedCategories), &categories)
	if err != nil {
		return nil, err
	}

	return categories, nil
}
