package controller

import (
	"jar-project/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

func CreateCategoryApi(c echo.Context) error {
	name := c.FormValue("name")
	description := c.FormValue("description")
	slug := c.FormValue("slug")
	result, err := model.CreateCategoryApi(name, description, slug)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, result)
}

func GetCategoryApi(c echo.Context) error {
	result, err := model.GetAllCategoriesApi()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"Message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, result)
}

func UpdateCategoryApi(c echo.Context) error {
	slug := c.Param("slug")
	name := c.FormValue("name")
	description := c.FormValue("description")
	result, err := model.UpdateCategoryApi(slug, name, description)
	if err != nil {
		return c.JSON(http.StatusOK, map[string]string{
			"Message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, result)
}

func DeleteCategoryApi(c echo.Context) error {
	slug := c.Param("slug")
	result, err := model.DeleteCategoryApi(slug)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"Message": err.Error(),
		})
	}
	return c.JSON(http.StatusNoContent, result)
}
