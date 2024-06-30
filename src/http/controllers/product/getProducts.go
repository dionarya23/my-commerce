package productv1controller

import (
	"net/http"

	"dionpamungkas.com/my-commerce/src/entities"
	productrepository "dionpamungkas.com/my-commerce/src/repositories/product"
	productusecase "dionpamungkas.com/my-commerce/src/usecase/product"
	"github.com/labstack/echo/v4"
)

func (i *V1Product) GetProducts(c echo.Context) (err error) {
	filters := &entities.ProductSearchFilter{}
	if categoryId := c.QueryParam("category_id"); categoryId != "" {
		if err != nil {
			return c.JSON(http.StatusBadRequest, ErrorResponse{
				Status:  false,
				Message: "Invalid value for 'category_id'",
			})
		}
		filters.CategoryId = categoryId
	}

	uu := productusecase.New(
		productrepository.New(i.DB),
		i.RedisClient,
	)

	data, err := uu.FindAll(filters)
	if err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Status:  false,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, SuccessResponse{
		Message: "Products found successfully",
		Data:    data,
	})
}
