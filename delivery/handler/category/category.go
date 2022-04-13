package catagory

import (
	"net/http"
	"project3/delivery/helper"
	_categoryUseCase "project3/usecase/category"

	"github.com/labstack/echo/v4"
)

type CategoryHandler struct {
	categoryUseCase _categoryUseCase.CategoryUseCaseInterface
}

func NewCategoryHandler(c _categoryUseCase.CategoryUseCaseInterface) CategoryHandler {
	return CategoryHandler{
		categoryUseCase: c,
	}
}

func (uh *CategoryHandler) GetAllCategoryHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		catagory, err := uh.categoryUseCase.GetAllCategory()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed to fetch data"))
		}

		responseCategories := []map[string]interface{}{}
		for i := 0; i < len(catagory); i++ {
			response := map[string]interface{}{
				"id":            catagory[i].ID,
				"catagory_name": catagory[i].CategoryName,
			}
			responseCategories = append(responseCategories, response)
		}

		return c.JSON(http.StatusOK, helper.ResponseSuccess("success get all catagories", responseCategories))
	}
}
