package v1routes

import (
	productv1controller "dionpamungkas.com/my-commerce/src/http/controllers/product"
	"dionpamungkas.com/my-commerce/src/http/middlewares"
)

func (i *V1Routes) MountProduct() {
	g := i.Echo.Group("/product")

	productController := productv1controller.New(&productv1controller.V1Product{
		DB:          i.DB,
		RedisClient: i.RedisClient,
	})

	g.GET("/", productController.GetProducts, middlewares.Authentication())
}
