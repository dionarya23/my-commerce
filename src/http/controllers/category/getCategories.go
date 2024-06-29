package categoryv1controller

import (
	"net/http"

	categoryrepository "dionpamungkas.com/my-commerce/src/repositories/category"
	categoryusecase "dionpamungkas.com/my-commerce/src/usecase/category"
	"github.com/labstack/echo/v4"
)

func (i *V1Category) GetCategories(c echo.Context) (err error) {
	uu := categoryusecase.New(
		categoryrepository.New(i.DB),
		i.RedisClient,
	)

	data, err := uu.FindAll()
	if err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Status:  false,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, SuccessResponse{
		Message: "Categories found successfully",
		Data:    data,
	})
}
