package v1routes

import (
	categoryv1controller "dionpamungkas.com/my-commerce/src/http/controllers/category"
	"dionpamungkas.com/my-commerce/src/http/middlewares"
)

func (i *V1Routes) MountCategory() {
	g := i.Echo.Group("/categories")

	categoryController := categoryv1controller.New(&categoryv1controller.V1Category{
		DB:          i.DB,
		RedisClient: i.RedisClient,
	})

	g.GET("/", categoryController.GetCategories, middlewares.Authentication())
}
